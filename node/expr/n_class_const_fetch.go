package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ClassConstFetch node
type ClassConstFetch struct {
	Meta         []meta.Meta
	Position     *position.Position
	Class        node.Node
	ConstantName node.Node
}

// NewClassConstFetch node constructor
func NewClassConstFetch(Class node.Node, ConstantName node.Node) *ClassConstFetch {
	return &ClassConstFetch{
		Class:        Class,
		ConstantName: ConstantName,
	}
}

// SetPosition sets node position
func (n *ClassConstFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClassConstFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *ClassConstFetch) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *ClassConstFetch) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *ClassConstFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassConstFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		v.EnterChildNode("Class", n)
		n.Class.Walk(v)
		v.LeaveChildNode("Class", n)
	}

	if n.ConstantName != nil {
		v.EnterChildNode("ConstantName", n)
		n.ConstantName.Walk(v)
		v.LeaveChildNode("ConstantName", n)
	}

	v.LeaveNode(n)
}
