package flow

import (
	"errors"
	"github.com/Megalotron/GoFlowMeter/packet"
	"io"
)

// BuildFromFile fills a flow Container using a PCAP file's content.
// It reads the file, extract the raw packet and converts it in a packet.Capsule.
// Then, the packet.Capsule is added in the Container for further computation.
func (c *Container) BuildFromFile(filename string) error {
	// Creates the file reader.
	reader, err := packet.NewFileReader(filename)
	if err != nil {
		return err
	}

	// For each packet... i.e. until we reach EOF.
	for i := uint64(0); ; i++ {
		// Extract the raw packet from PCAP file.
		rawPacket, err := reader.GetNextPacket()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		// Convert raw packet to capsule.
		capsule, err := packet.NewCapsuleFromPCAPPacket(rawPacket, i)
		if err != nil {
			if errors.Is(err, packet.ErrInvalidPacket) {
				i -= 1

				continue
			}

			return err
		}

		// Adds the capsule to the current flow container.
		c.capsules = append(c.capsules, *capsule)
	}

	return nil
}
