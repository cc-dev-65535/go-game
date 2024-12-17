package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"game/test/images"

	"github.com/Kangaroux/go-spritesheet"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

type Game struct {
	state     string
	positionX int
	positionY int
}

const tileSize = 16
const mapPath = "maps/game.tmx"

var (
	characterImage   *ebiten.Image
	characterSprites map[string]*spritesheet.Sprite
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Character_png))
	if err != nil {
		panic(err)
	}
	characterImage = ebiten.NewImageFromImage(img)

	characterSheet, err := spritesheet.OpenAndRead("character.yaml")
	if err != nil {
		panic(err)
	}
	characterSprites = characterSheet.Sprites()
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.state = "left_1"
		g.positionX -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.state = "right_1"
		g.positionX += 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.state = "up_1"
		g.positionY -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.state = "down_1"
		g.positionY += 1
	}
	return nil
}

func getImageFromSpritesheet(imageFile *ebiten.Image, sprite string) *ebiten.Image {
	return imageFile.SubImage(characterSprites[sprite].Rect()).(*ebiten.Image)
}

func (g *Game) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, 0)

	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		panic(err)
	}

	// fmt.Println(gameMap)

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

	var layerOneImage image.Image = renderer.Result

	renderer.Clear()

	err = renderer.RenderLayer(1)
	if err != nil {
		fmt.Printf("layer unsupported for rendering: %s", err.Error())
		panic(err)
	}

	var layerTwoImage image.Image = renderer.Result

	renderer.Clear()

	screen.DrawImage(ebiten.NewImageFromImage(layerOneImage), options)
	screen.DrawImage(ebiten.NewImageFromImage(layerTwoImage), options)

	optionsCharacter := &ebiten.DrawImageOptions{}
	optionsCharacter.GeoM.Translate(float64((g.positionX)*tileSize), float64((g.positionY)*tileSize))

	screen.DrawImage(getImageFromSpritesheet(characterImage, g.state), optionsCharacter)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tileSize * 45, tileSize * 28
}

func main() {
	game := &Game{state: "down_1", positionX: 1, positionY: 1}
	ebiten.SetTPS(10)
	ebiten.SetWindowSize(tileSize*45*2, tileSize*28*2)
	ebiten.SetWindowTitle("Top Down World")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
