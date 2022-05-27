package packet

import (
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"strconv"
)

var (
	ErrInvalidPacket = errors.New("invalid or corrupted packet")
)

// Capsule contains some of the packet's fields.
// Using multiple of them, we can create a Flow.
type Capsule struct {
	// The capsule's identifier.
	id uint64

	// The packet's source address.
	srcAddress string
	// The packet's destination address.
	dstAddress string

	// The packet's source port.
	srcPort uint16
	// The packet's destination port.
	dstPort uint16

	// The protocol number referencing the protocol used by the packet to transit.
	protocol ProtocolNumber
}

// NewCapsuleFromPacket creates a new Capsule using a gopacket.Packet to extract relevant fields.
func NewCapsuleFromPacket(packet gopacket.Packet, id uint64) (*Capsule, error) {
	// The packet misses crucial information.
	if packet.NetworkLayer() == nil || packet.TransportLayer() == nil {
		return nil, ErrInvalidPacket
	}

	// Packet's source and destination addresses.
	srcAddress, dstAddress := packet.NetworkLayer().NetworkFlow().Endpoints()

	// Packet's source and destination ports.
	srcPortRaw, dstPortRaw := packet.TransportLayer().TransportFlow().Endpoints()

	srcPort, err := strconv.ParseUint(srcPortRaw.String(), 10, 16)
	if err != nil {
		return nil, err
	}

	dstPort, err := strconv.ParseUint(dstPortRaw.String(), 10, 16)
	if err != nil {
		return nil, err
	}

	// Packet's protocol number.
	protocol := ProtocolUDP
	if packet.Layer(layers.LayerTypeTCP) != nil {
		protocol = ProtocolTCP
	}

	// Create the capsule.
	return &Capsule{
		id:         id,
		srcAddress: srcAddress.String(),
		dstAddress: dstAddress.String(),
		srcPort:    uint16(srcPort),
		dstPort:    uint16(dstPort),
		protocol:   protocol,
	}, nil
}
