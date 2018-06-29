package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Expression node
type Expression struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewExpression node constructor
func NewExpression(Expr node.Node) *Expression {
	return &Expression{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Expression) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Expression) GetPosition() *position.Position {
	return n.Position
}

func (n *Expression) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Expression) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Expression) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Expression) Walk(v walker.Visitor) {
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
