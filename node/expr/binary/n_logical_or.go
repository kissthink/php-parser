package binary

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// LogicalOr node
type LogicalOr struct {
	Meta     []meta.Meta
	Position *position.Position
	Left     node.Node
	Right    node.Node
}

// NewLogicalOr node constructor
func NewLogicalOr(Variable node.Node, Expression node.Node) *LogicalOr {
	return &LogicalOr{
		Left:  Variable,
		Right: Expression,
	}
}

// SetPosition sets node position
func (n *LogicalOr) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *LogicalOr) GetPosition() *position.Position {
	return n.Position
}

func (n *LogicalOr) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *LogicalOr) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *LogicalOr) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *LogicalOr) Walk(v walker.Visitor) {
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
