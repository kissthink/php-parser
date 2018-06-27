package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/pkg/profile"
	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
)

var wg sync.WaitGroup
var usePhp5 *bool
var dumpType string
var profiler string
var showComments *bool
var showResolvedNs *bool

var toPrint []string

func main() {
	usePhp5 = flag.Bool("php5", false, "parse as PHP5")
	showComments = flag.Bool("c", false, "show comments")
	showResolvedNs = flag.Bool("r", false, "resolve names")
	flag.StringVar(&dumpType, "d", "", "dump format: [custom, go, json, pretty_json]")
	flag.StringVar(&profiler, "prof", "", "start profiler: [cpu, mem]")

	flag.Parse()

	switch profiler {
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()
	}

	pathCh := make(chan []byte)
	resultCh := make(chan parser.Parser)

	// run 4 concurrent parserWorkers
	for i := 0; i < 10; i++ {
		go parserWorker(pathCh, resultCh)
	}

	// run printer goroutine
	go printer(resultCh)

	// process files
	processPath(flag.Args(), pathCh)

	// wait the all files done
	wg.Wait()
	close(pathCh)
	close(resultCh)

	for _, str := range toPrint {
		io.WriteString(os.Stdout, str+"\n")
	}
}

func processPath(pathList []string, pathCh chan<- []byte) {
	var ps [][]byte

	for _, path := range pathList {
		real, err := realpath.Realpath(path)
		checkErr(err)

		s, err := os.Stat(real)
		checkErr(err)

		if !s.IsDir() {
			wg.Add(1)
			c, err := ioutil.ReadFile(real)
			checkErr(err)
			ps = append(ps, c)
		} else {
			godirwalk.Walk(real, &godirwalk.Options{
				Unsorted: true,
				Callback: func(osPathname string, de *godirwalk.Dirent) error {
					if !de.IsDir() && filepath.Ext(osPathname) == ".php" {
						wg.Add(1)
						c, err := ioutil.ReadFile(osPathname)
						checkErr(err)
						ps = append(ps, c)
					}
					return nil
				},
				ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
					return godirwalk.SkipNode
				},
			})
		}
	}

	for _, p := range ps {
		pathCh <- p
	}
}

func parserWorker(pathCh <-chan []byte, result chan<- parser.Parser) {
	var parserWorker parser.Parser

	for {
		path, ok := <-pathCh
		if !ok {
			return
		}

		src := bytes.NewReader(path)

		if *usePhp5 {
			parserWorker = php5.NewParser(src, "")
		} else {
			parserWorker = php7.NewParser(src, "")
		}

		parserWorker.Parse()

		result <- parserWorker
	}
}

func printer(result <-chan parser.Parser) {
	var counter int

	for {
		parserWorker, ok := <-result
		if !ok {
			return
		}

		counter++

		toPrint = append(toPrint, "==> ["+strconv.Itoa(counter)+"]")

		if dumpType == "" {
			wg.Done()
			continue
		}

		for _, e := range parserWorker.GetErrors() {
			fmt.Println(e)
		}

		var nsResolver *visitor.NamespaceResolver
		if *showResolvedNs {
			nsResolver = visitor.NewNamespaceResolver()
			parserWorker.GetRootNode().Walk(nsResolver)
		}

		var comments parser.Comments
		if *showComments {
			comments = parserWorker.GetComments()
		}

		switch dumpType {
		case "custom":
			dumper := &visitor.Dumper{
				Writer:     os.Stdout,
				Indent:     "| ",
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "json":
			dumper := &visitor.JsonDumper{
				Writer:     os.Stdout,
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "pretty_json":
			dumper := &visitor.PrettyJsonDumper{
				Writer:     os.Stdout,
				Comments:   comments,
				NsResolver: nsResolver,
			}
			parserWorker.GetRootNode().Walk(dumper)
		case "go":
			dumper := &visitor.GoDumper{Writer: os.Stdout}
			parserWorker.GetRootNode().Walk(dumper)
		}

		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
