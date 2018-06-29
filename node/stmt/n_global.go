package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Global node
type Global struct {
	Meta     []meta.Meta
	Position *position.Position
	Vars     []node.Node
}

// NewGlobal node constructor
func NewGlobal(Vars []node.Node) *Global {
	return &Global{
		Vars: Vars,
	}
}

// SetPosition sets node position
func (n *Global) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Global) GetPosition() *position.Position {
	return n.Position
}

func (n *Global) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Global) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Global) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Global) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		v.EnterChildList("Vars", n)
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Vars", n)
	}

	v.LeaveNode(n)
}
