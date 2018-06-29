package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Echo node
type Echo struct {
	Meta     []meta.Meta
	Position *position.Position
	Exprs    []node.Node
}

// NewEcho node constructor
func NewEcho(Exprs []node.Node) *Echo {
	return &Echo{
		Exprs: Exprs,
	}
}

// SetPosition sets node position
func (n *Echo) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Echo) GetPosition() *position.Position {
	return n.Position
}

func (n *Echo) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Echo) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Echo) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Echo) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		v.EnterChildList("Exprs", n)
		for _, nn := range n.Exprs {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Exprs", n)
	}

	v.LeaveNode(n)
}
