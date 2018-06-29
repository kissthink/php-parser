package assign

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Mul node
type Mul struct {
	Meta       []meta.Meta
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewMul node constructor
func NewMul(Variable node.Node, Expression node.Node) *Mul {
	return &Mul{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *Mul) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Mul) GetPosition() *position.Position {
	return n.Position
}

func (n *Mul) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Mul) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Mul) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Mul) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Expression != nil {
		v.EnterChildNode("Expression", n)
		n.Expression.Walk(v)
		v.LeaveChildNode("Expression", n)
	}

	v.LeaveNode(n)
}
