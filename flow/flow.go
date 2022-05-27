package flow

import (
	"github.com/Megalotron/GoFlowMeter/packet"
)

// Container represents a flow.
type Container struct {
	capsules []packet.Capsule
}

// NewContainer creates an empty container for capsules.
func NewContainer() *Container {
	return &Container{}
}

// addCapsule contains adds a valid capsule in the current flow Container.
// TODO: Add more details when the method is finished.
func (c *Container) addCapsule(capsule *packet.Capsule) {
	c.capsules = append(c.capsules, *capsule)
}
