// Package gosysfonts gets system TTF fonts by name
package gosysfonts

import (
	"github.com/golang/freetype/truetype"
)

// Pool is a pool of system fonts
type Pool interface {
	GetFont(name string) (*truetype.Font, error)
}
