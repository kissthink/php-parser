package expr

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// PostInc node
type PostInc struct {
	Meta     []meta.Meta
	Position *position.Position
	Variable node.Node
}

// NewPostInc node constructor
func NewPostInc(Variable node.Node) *PostInc {
	return &PostInc{
		Variable: Variable,
	}
}

// SetPosition sets node position
func (n *PostInc) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *PostInc) GetPosition() *position.Position {
	return n.Position
}

func (n *PostInc) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *PostInc) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *PostInc) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PostInc) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	v.LeaveNode(n)
}
