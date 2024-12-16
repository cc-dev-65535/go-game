package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"game/test/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
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
	tileXCount := width / tileSize

	sx := (cell % tileXCount) * tileSize
	sy := (cell / tileXCount) * tileSize

	return tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)

	gameMap, err := tiled.LoadFile("maps/game.tmx")
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		panic(err)
	}

	fmt.Println(gameMap)

	// You can also render the map to an in-memory image for direct
	// use with the default Renderer, or by making your own.
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Render just layer 0 to the Renderer.
	err = renderer.RenderLayer(0)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		os.Exit(2)
	}

	// Get a reference to the Renderer's output, an image.NRGBA struct.
	img := renderer.Result
	var testImage image.Image = img

	// make image
	// ebiten.NewImageFromImage(testImage)

	// Clear the render result after copying the output if separation of
	// layers is desired.
	renderer.Clear()

	// And so on. You can also export the image to a file by using the
	// Renderer's Save functions.

	screen.DrawImage(ebiten.NewImageFromImage(testImage), op)

	// const xCount = 320 / tileSize
	// for _, layer := range g.layers {
	// 	for i, cell := range layer {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

	// 		screen.DrawImage(getImageFromSpritesheet(tilesImage, cell), op)
	// 	}
	// }
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
