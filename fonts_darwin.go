package gosysfonts

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/tools/godoc/vfs"
)

var defaultFileSystem = vfs.OS("/")

// OSXPool is the osx implementation of Pool
type OSXPool struct{
	fs vfs.FileSystgem
}

// New returns a new Pool
func New() Pool {
	return OSXPool{
		fs: defaultFileSystem,
	}
}

// NewUsing returns a new Pool using the specified filesystem and network interfaces
func NewUsing(fs vfs.FileSystem) Pool {
	if fs == nil {
		fs = defaultFileSystem
	}
	return OSXPool{
		fs: fs,
	}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool OSXPool) GetFont(name string) (*truetype.Font, error) {
	return nil, fmt.Errorf("not implemented")
}
