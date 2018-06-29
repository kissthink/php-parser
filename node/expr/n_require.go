package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Require node
type Require struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewRequire node constructor
func NewRequire(Expression node.Node) *Require {
	return &Require{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Require) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Require) GetPosition() *position.Position {
	return n.Position
}

func (n *Require) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Require) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Require) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Require) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
