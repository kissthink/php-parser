package binary

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Equal node
type Equal struct {
	Meta     []meta.Meta
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewEqual node constructor
func NewEqual(Variable node.Node, Expression node.Node) *Equal {
	return &Equal{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *Equal) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Equal) GetPosition() *position.Position {
	return n.Position
}

func (n *Equal) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Equal) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Equal) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Equal) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
