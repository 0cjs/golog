package gelf

import (
	"testing"
)

type tsWriter struct{}

func (tsWriter) Write([]byte) (int, error) { return 0, nil }

func TestNewOnStream(t *testing.T) {
	_ = NewOnStream(&tsWriter{}, 0)    // Delimiter is "null char"
	_ = NewOnStream(&tsWriter{}, '\n') // Delimiter is newline
}
