package gosysfonts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Equal(t, LinuxPool{}, New())
}

func TestGetFont(t *testing.T) {
	font, err := New().GetFont("anything")
	assert.Nil(t, font)
	assert.EqualError(t, err, "not implemented")
}
