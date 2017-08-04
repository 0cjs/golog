package gelf

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	assert := assert.New(t)

	msg := "Help me!\nI've fallen and I can't get up!"
	var m Message = NewMessage(msg)

	assert.NotEmpty(m.Host)
	assert.Equal(msg, m.ShortMessage)
	assert.Equal("", m.FullMessage)
	// XXX assert.NotEmpty(m.Timestamp)
	assert.EqualValues(LOG_NOTICE, m.Level)
	assert.Empty(m.Extra)
}
