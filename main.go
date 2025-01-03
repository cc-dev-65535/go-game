package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"strings"

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
const tileLayers = 3

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
		if g.state == "left_3" {
			g.state = "left_4"
		} else {
			g.state = "left_3"
		}
		g.positionX -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if g.state == "right_3" {
			g.state = "right_4"
		} else {
			g.state = "right_3"
		}
		g.positionX += 1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if g.state == "up_3" {
			g.state = "up_4"
		} else {
			g.state = "up_3"
		}
		g.positionY -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if g.state == "down_3" {
			g.state = "down_4"
		} else {
			g.state = "down_3"
		}
		g.positionY += 1
	} else if g.state != "" {
		result := strings.Split(g.state, "_")
		if g.state == result[0]+"_1" {
			g.state = result[0] + "_2"
		} else {
			g.state = result[0] + "_1"
		}
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

	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		fmt.Printf("map unsupported for rendering: %s", err.Error())
		panic(err)
	}

	for i := 0; i < tileLayers; i++ {
		err = renderer.RenderLayer(i)
		if err != nil {
			fmt.Printf("layer unsupported for rendering: %s", err.Error())
			panic(err)
		}

		var layerImage image.Image = renderer.Result

		screen.DrawImage(ebiten.NewImageFromImage(layerImage), options)

		renderer.Clear()
	}

	optionsCharacter := &ebiten.DrawImageOptions{}
	optionsCharacter.GeoM.Translate(float64((g.positionX)*tileSize), float64((g.positionY)*tileSize))

	screen.DrawImage(getImageFromSpritesheet(characterImage, g.state), optionsCharacter)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return tileSize * 45, tileSize * 28
}

func main() {
	game := &Game{state: "down_1", positionX: 1, positionY: 1}
	ebiten.SetTPS(8)
	ebiten.SetWindowSize(tileSize*45*2, tileSize*28*2)
	ebiten.SetWindowTitle("Top-Down World")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
