package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Eval node
type Eval struct {
	Meta     []meta.Meta
	Position *position.Position
	Expr     node.Node
}

// NewEval node constructor
func NewEval(Expression node.Node) *Eval {
	return &Eval{
		Expr: Expression,
	}
}

// SetPosition sets node position
func (n *Eval) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Eval) GetPosition() *position.Position {
	return n.Position
}

func (n *Eval) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Eval) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Eval) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Eval) Walk(v walker.Visitor) {
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
