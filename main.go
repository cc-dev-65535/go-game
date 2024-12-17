package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"game/test/images"

	"github.com/Kangaroux/go-spritesheet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type Game struct{ player []int }

const tileSize = 16
const mapPath = "maps/game.tmx"

var (
	characterImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Character_png))
	if err != nil {
		log.Fatal(err)
	}
	characterImage = ebiten.NewImageFromImage(img)

	sheet, err := spritesheet.OpenAndRead("spritesheet.yml")
}

func (g *Game) Update() error {
	return nil
}

func getImageFromSpritesheet(imageFile *ebiten.Image, cell int) *ebiten.Image {
	width := imageFile.Bounds().Dx()
	tileXCount := width / tileSize

	sx := (cell % tileXCount) * tileSize
	sy := (cell / tileXCount) * tileSize

	sy = 16
	fmt.Println(sx)
	fmt.Println(sy)
	return imageFile.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image)
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, 0)

	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		panic(err)
	}

	fmt.Println(gameMap)

	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		panic(err)
	}

	err = renderer.RenderLayer(0)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		panic(err)
	}

	var layerImage image.Image = renderer.Result

	renderer.Clear()

	screen.DrawImage(ebiten.NewImageFromImage(layerImage), options)

	optionsCharacter := &ebiten.DrawImageOptions{}
	optionsCharacter.GeoM.Translate(float64((0)*tileSize), float64((0)*tileSize))

	screen.DrawImage(getImageFromSpritesheet(characterImage, 1), optionsCharacter)

	// const xCount = 320 / tileSize
	// for _, layer := range g.layers {
	// 	for i, cell := range layer {
	// 		op := &ebiten.DrawImageOptions{}
	// 		op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

	// 		screen.DrawImage(getImageFromSpritesheet(characterImage, cell), op)
	// 	}
	// }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tileSize * 45, tileSize * 28
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(tileSize*45*2, tileSize*28*2)
	ebiten.SetWindowTitle("My Game")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
