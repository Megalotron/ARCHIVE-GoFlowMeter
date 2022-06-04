package flow

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		name               string
		filename           string
		expectedErr        error
		expectedNumCapsule int
	}{
		{
			name:               "Ethernet PCAP",
			filename:           "../test_ethernet.pcap",
			expectedErr:        nil,
			expectedNumCapsule: 10,
		},
		{
			name:               "Random PCAP",
			filename:           "../test_various.pcap",
			expectedErr:        nil,
			expectedNumCapsule: 1867,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			container := NewContainer()

			err := container.BuildFromFile(tt.filename)
			assert.Equal(t, tt.expectedErr, err)

			if tt.expectedErr == nil {
				fmt.Println(len(container.capsules))
				assert.Len(t, container.capsules, tt.expectedNumCapsule)
			}
		})
	}
}
