package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ArrayItem node
type ArrayItem struct {
	Meta     []meta.Meta
	Position *position.Position
	Key      node.Node
	Val      node.Node
}

// NewArrayItem node constructor
func NewArrayItem(Key node.Node, Val node.Node) *ArrayItem {
	return &ArrayItem{
		Key: Key,
		Val: Val,
	}
}

// SetPosition sets node position
func (n *ArrayItem) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ArrayItem) GetPosition() *position.Position {
	return n.Position
}

func (n *ArrayItem) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ArrayItem) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ArrayItem) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ArrayItem) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Key != nil {
		v.EnterChildNode("Key", n)
		n.Key.Walk(v)
		v.LeaveChildNode("Key", n)
	}

	if n.Val != nil {
		v.EnterChildNode("Val", n)
		n.Val.Walk(v)
		v.LeaveChildNode("Val", n)
	}

	v.LeaveNode(n)
}
