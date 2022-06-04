package packet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadValidFile(t *testing.T) {
	filename := "../test_ethernet.pcap"

	reader, err := NewFileReader(filename)
	assert.NoError(t, err)

	reader.Close()
}
