package stmt

import (
	"github.com/z7zmey/php-parser/meta"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Static node
type Static struct {
	Meta     []meta.Meta
	Position *position.Position
	Vars     []node.Node
}

// NewStatic node constructor
func NewStatic(Vars []node.Node) *Static {
	return &Static{
		Vars: Vars,
	}
}

// SetPosition sets node position
func (n *Static) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Static) GetPosition() *position.Position {
	return n.Position
}

func (n *Static) AddMeta(m []meta.Meta) {
	n.Meta = append(n.Meta, m...)
}

func (n *Static) GetMeta() []meta.Meta {
	return n.Meta
}

// Attributes returns node attributes as map
func (n *Static) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Static) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		v.EnterChildList("Vars", n)
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Vars", n)
	}

	v.LeaveNode(n)
}
