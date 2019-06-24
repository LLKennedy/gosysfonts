package gosysfonts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/godoc/vfs"
)

func TestNew(t *testing.T) {
	assert.Equal(t, WinPool{fs: defaultFileSystem}, New())
}

func TestNewUsing(t *testing.T) {
	assert.Equal(t, WinPool{fs: defaultFileSystem}, NewUsing(nil))
	assert.Equal(t, WinPool{fs: vfs.OS(".")}, NewUsing(vfs.OS(".")))
}

func TestGetFont(t *testing.T) {
	t.Run("invalid file", func(t *testing.T) {
		font, err := New().GetFont("gibberish font that doesn't exist")
		assert.Nil(t, font)
		assert.EqualError(t, err, "open C:/Windows/Fonts/gibberish font that doesn't exist.ttf: The system cannot find the file specified.")
	})
	t.Run("valid font", func(t *testing.T) {
		_, err := New().GetFont("arial")
		assert.NoError(t, err)
	})
}
