package cast

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Double node
type Double struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewDouble node constructor
func NewDouble(Expr node.Node) *Double {
	return &Double{
		Expr: Expr,
	}
}

// SetPosition sets node position
func (n *Double) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Double) GetPosition() *position.Position {
	return n.Position
}

func (n *Double) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Double) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Double) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Double) Walk(v walker.Visitor) {
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
