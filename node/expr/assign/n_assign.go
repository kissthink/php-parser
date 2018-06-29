package assign

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Assign node
type Assign struct {
	Meta       []meta.Meta
	Position   *position.Position
	Variable   node.Node
	Expression node.Node
}

// NewAssign node constructor
func NewAssign(Variable node.Node, Expression node.Node) *Assign {
	return &Assign{
		Variable:   Variable,
		Expression: Expression,
	}
}

// SetPosition sets node position
func (n *Assign) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Assign) GetPosition() *position.Position {
	return n.Position
}

func (n *Assign) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Assign) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Assign) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Assign) Walk(v walker.Visitor) {
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
