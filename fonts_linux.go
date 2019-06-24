package gosysfonts

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"golang.org/x/tools/godoc/vfs"
)

var defaultFileSystem = vfs.OS("/")

// LinuxPool is the linux implementation of Pool
type LinuxPool struct{
	fs vfs.FileSystem
}

// New returns a new Pool
func New() Pool {
	return LinuxPool{
		fs: defaultFileSystem,
	}
}

// NewUsing returns a new Pool using the specified filesystem and network interfaces
func NewUsing(fs vfs.FileSystem) Pool {
	if fs == nil {
		fs = defaultFileSystem
	}
	return LinuxPool{
		fs: fs,
	}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool LinuxPool) GetFont(name string) (*truetype.Font, error) {
	return nil, fmt.Errorf("not implemented")
}
