// +build windows

package gosysfonts

import (
	"github.com/golang/freetype/truetype"
	"io/ioutil"
)

type WinPool struct{}

func New() Pool {
	return WinPool{}
}

func (pool WinPool) GetFont(name string) (*truetype.Font, error) {
	ttfData, err := ioutil.ReadFile("%WINDIR%/Fonts/calibri.ttf")
	if err != nil {
		return nil, err
	}
	return truetype.Parse(ttfData)
}
