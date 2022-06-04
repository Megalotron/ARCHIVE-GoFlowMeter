package packet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCapsuleFromPacket(t *testing.T) {
	testCases := []struct {
		name       string
		filename   string
		err        error
		id         uint64
		srcAddress string
		dstAddress string
		srcPort    uint16
		dstPort    uint16
		protocol   ProtocolNumber
	}{
		{
			name:       "Valid Local Ethernet Connection",
			filename:   "../test_ethernet.pcap",
			err:        nil,
			id:         0,
			srcAddress: "10.1.1.2",
			dstAddress: "10.1.1.1",
			srcPort:    44644,
			dstPort:    80,
			protocol:   ProtocolTCP,
		},
		{
			name:       "Valid Spotify Response",
			filename:   "../test_various.pcap",
			err:        nil,
			id:         0,
			srcAddress: "20.82.247.128",
			dstAddress: "192.168.2.163",
			srcPort:    443,
			dstPort:    12425,
			protocol:   ProtocolTCP,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := NewFileReader(tt.filename)
			assert.NoError(t, err)

			packet, err := reader.GetNextPacket()
			assert.NoError(t, err)

			capsule, err := NewCapsuleFromPCAPPacket(packet, tt.id)
			assert.Equal(t, tt.err, err)

			if tt.err == nil {
				assert.Equal(t, tt.id, capsule.id)
				assert.Equal(t, tt.srcAddress, capsule.srcAddress)
				assert.Equal(t, tt.dstAddress, capsule.dstAddress)
				assert.Equal(t, tt.srcPort, capsule.srcPort)
				assert.Equal(t, tt.dstPort, capsule.dstPort)
				assert.Equal(t, tt.protocol, capsule.protocol)
			}

			reader.Close()
		})
	}
}
