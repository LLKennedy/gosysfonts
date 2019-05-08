// +build windows

package gosysfonts

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"fmt"
)

// WinPool is the Windows implementation of Pool
type WinPool struct{}

// New returns a new Pool
func New() Pool {
	return WinPool{}
}

// GetFont returns the truetype font corresponding to the font name passed in
func (pool WinPool) GetFont(name string) (*truetype.Font, error) {
	ttfData, err := ioutil.ReadFile(fmt.Sprintf("C:/Windows/Fonts/%s.ttf", name))
	if err != nil {
		return nil, err
	}
	return truetype.Parse(ttfData)
}
