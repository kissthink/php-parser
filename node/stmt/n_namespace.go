package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Namespace node
type Namespace struct {
	Meta          []meta.Meta
	Position      *position.Position
	NamespaceName node.Node
	Stmts         []node.Node
}

// NewNamespace node constructor
func NewNamespace(NamespaceName node.Node, Stmts []node.Node) *Namespace {
	return &Namespace{
		NamespaceName: NamespaceName,
		Stmts:         Stmts,
	}
}

// SetPosition sets node position
func (n *Namespace) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Namespace) GetPosition() *position.Position {
	return n.Position
}

func (n *Namespace) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Namespace) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Namespace) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Namespace) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.NamespaceName != nil {
		v.EnterChildNode("NamespaceName", n)
		n.NamespaceName.Walk(v)
		v.LeaveChildNode("NamespaceName", n)
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
