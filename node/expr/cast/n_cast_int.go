package cast

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Int node
type Int struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewInt node constructor
func NewInt(Expr node.Node) *Int {
	return &Int{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Int) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Int) GetPosition() *position.Position {
	return n.Position
}

func (n *Int) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Int) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Int) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Int) Walk(v walker.Visitor) {
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
