package gosysfonts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/godoc/vfs"
)

func TestNew(t *testing.T) {
	assert.Equal(t, OSXPool{fs: defaultFileSystem}, New())
}

func TestNewUsing(t *testing.T) {
	assert.Equal(t, OSXPool{fs: defaultFileSystem}, NewUsing(nil))
	assert.Equal(t, OSXPool{fs: vfs.OS(".")}, NewUsing(vfs.OS(".")))
}

func TestGetFont(t *testing.T) {
	font, err := New().GetFont("anything")
	assert.Nil(t, font)
	assert.EqualError(t, err, "not implemented")
}
