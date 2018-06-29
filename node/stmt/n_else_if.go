package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ElseIf node
type ElseIf struct {
	Meta     []meta.Meta
	Position *position.Position
	Cond     node.Node
	Stmt     node.Node
}

// NewElseIf node constructor
func NewElseIf(Cond node.Node, Stmt node.Node) *ElseIf {
	return &ElseIf{
		Cond: Cond,
		Stmt: Stmt,
	}
}

// SetPosition sets node position
func (n *ElseIf) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ElseIf) GetPosition() *position.Position {
	return n.Position
}

func (n *ElseIf) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ElseIf) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ElseIf) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ElseIf) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
