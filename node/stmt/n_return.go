package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Return node
type Return struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewReturn node constructor
func NewReturn(Expr node.Node) *Return {
	return &Return{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Return) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Return) GetPosition() *position.Position {
	return n.Position
}

func (n *Return) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Return) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Return) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Return) Walk(v walker.Visitor) {
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
