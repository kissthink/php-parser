package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// New node
type New struct {
	Meta         []meta.Meta
	Position     *position.Position
	Class        node.Node
	ArgumentList *node.ArgumentList
}

// NewNew node constructor
func NewNew(Class node.Node, ArgumentList *node.ArgumentList) *New {
	return &New{
		Class:        Class,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *New) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *New) GetPosition() *position.Position {
	return n.Position
}

func (n *New) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *New) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *New) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *New) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		v.EnterChildNode("Class", n)
		n.Class.Walk(v)
		v.LeaveChildNode("Class", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	v.LeaveNode(n)
}
