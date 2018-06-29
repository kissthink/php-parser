package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// StmtList node
type StmtList struct {
	Meta     []meta.Meta
	Position *position.Position
	Stmts    []node.Node
}

// NewStmtList node constructor
func NewStmtList(Stmts []node.Node) *StmtList {
	return &StmtList{
		Stmts: Stmts,
	}
}

// SetPosition sets node position
func (n *StmtList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *StmtList) GetPosition() *position.Position {
	return n.Position
}

func (n *StmtList) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *StmtList) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *StmtList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StmtList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
