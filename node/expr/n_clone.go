package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Clone node
type Clone struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewClone node constructor
func NewClone(Expression node.Node) *Clone {
	return &Clone{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Clone) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Clone) GetPosition() *position.Position {
	return n.Position
}

func (n *Clone) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Clone) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Clone) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Clone) Walk(v walker.Visitor) {
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
