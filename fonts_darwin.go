package gosysfonts

import (
	"fmt"
	"github.com/golang/freetype/truetype"
)

// OSXPool is the osx implementation of Pool
type OSXPool struct{}

// New returns a new Pool
func New() Pool {
	return OSXPool{}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool OSXPool) GetFont(name string) (*truetype.Font, error) {
	return nil, fmt.Errorf("not implemented")
}
