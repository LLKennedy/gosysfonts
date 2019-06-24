package gosysfonts

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/tools/godoc/vfs"
)

var defaultFileSystem = vfs.OS("C:/")

// WinPool is the Windows implementation of Pool
type WinPool struct {
	fs vfs.FileSystem
}

// New returns a new Pool
func New() Pool {
	return WinPool{
		fs: defaultFileSystem,
	}
}

// NewUsing returns a new Pool using the specified filesystem and network interfaces
func NewUsing(fs vfs.FileSystem) Pool {
	if fs == nil {
		fs = defaultFileSystem
	}
	return WinPool{
		fs: fs,
	}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool WinPool) GetFont(name string) (*truetype.Font, error) {
	ttfData, err := ioutil.ReadFile(fmt.Sprintf("C:/Windows/Fonts/%s.ttf", name))
	if err != nil {
		return nil, err
	}
	return truetype.Parse(ttfData)
}
