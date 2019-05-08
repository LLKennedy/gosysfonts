// +build linux

package gosysfonts

import (
	"github.com/golang/freetype/truetype"
	"fmt"
)

// LinuxPool is the linux implementation of Pool
type LinuxPool struct{}

// New returns a new Pool
func New() Pool {
	return LinuxPool{}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool LinuxPool) GetFont(name string) (*truetype.Font, error) {
	return nil, fmt.Errorf("not implemented")
}
