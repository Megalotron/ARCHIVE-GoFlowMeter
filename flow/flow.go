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
