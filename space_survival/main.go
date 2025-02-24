package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	player Player
	// laser Laser
}


type Player struct{
	character *ebiten.Image
	rad float64
	x float64
	y float64
	dirX float64
	dirY float64
	height int
	width int
}

func (p *Player) createCharacter() {
	p.character = ebiten.NewImage(p.height, p.width)
}

// type Laser struct{

// }

func (g *Game) PlayerMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.rad -= 0.1
		g.player.dirX = math.Sin(g.player.rad)
		g.player.dirY = math.Cos(g.player.rad)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.rad += 0.1
		g.player.dirX = math.Sin(g.player.rad)
		g.player.dirY = math.Cos(g.player.rad)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.y -= g.player.dirY
		g.player.x += g.player.dirX
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.y += g.player.dirY
		g.player.x -= g.player.dirX
	}
}

func (g *Game) PlayerShoot() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {

	}
}

func (g *Game) Update() error {
	g.PlayerMovement()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	geo := ebiten.GeoM{}
	geo.Translate(-5, -5)
	geo.Rotate(p.rad)

	geo.Translate(p.x, p.y)
	op := &ebiten.DrawImageOptions{GeoM:geo}
	screen.DrawImage(p.character, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Ticks Per Second: %0.2f", ebiten.ActualTPS()))
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	player := Player{
		height : 10,
		width : 10,
		x : 475,
		y : 265,
	}
	player.createCharacter()
	player.character.Fill(color.White)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)
	if err := ebiten.RunGame(&Game{player}); err != nil {
		log.Fatal(err)
	}
	

}
