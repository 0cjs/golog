package gelf

import (
	"testing"
)

type psWriter struct{}

func (psWriter) Write([]byte) (int, error) { return 0, nil }

func TestNewOnPacket(t *testing.T) {
	_, _ = NewOnPacket(&psWriter{}, DefaultMTU, None, 0)
	_, _ = NewOnPacket(&psWriter{}, 1024*1024, Gzip, DefaultCompressionLevel)
	// XXX Test that MTU < ??? fails
}
