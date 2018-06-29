package binary

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Greater node
type Greater struct {
	Meta     []meta.Meta
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewGreater node constructor
func NewGreater(Variable node.Node, Expression node.Node) *Greater {
	return &Greater{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *Greater) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Greater) GetPosition() *position.Position {
	return n.Position
}

func (n *Greater) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Greater) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Greater) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Greater) Walk(v walker.Visitor) {
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
