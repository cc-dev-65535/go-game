package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"game/test/images"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{ layers [][]int }

const (
	tileSize = 16
)

var (
	tilesImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Grass_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	return nil
}

func getImageFromSpritesheet(imageFile *ebiten.Image, cell int) *ebiten.Image {
	width := imageFile.Bounds().Dx()
	fmt.Println(width)
	tileXCount := width / tileSize
	fmt.Println(tileXCount)

	sx := (cell % tileXCount) * tileSize
	sy := (cell / tileXCount) * tileSize

	return tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)

	const xCount = 320 / tileSize
	for _, layer := range g.layers {
		for i, cell := range layer {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			screen.DrawImage(getImageFromSpritesheet(tilesImage, cell), op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{layers: layers}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("My Game")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
