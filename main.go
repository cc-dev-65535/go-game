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

type Game struct{}

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
	fmt.Println("hi")
	tilesImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	fmt.Println(w)
	tileXCount := w / tileSize
	fmt.Println(tileXCount)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	fmt.Println("hijhg")
	screen.DrawImage(tilesImage.SubImage(image.Rect(0, 0, 50, 50)).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("My Game")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
