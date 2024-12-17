package images

import (
	_ "embed"
)

var (
	//go:embed Grass.png
	Grass_png []byte

	//go:embed Character.png
	Character_png []byte
)
