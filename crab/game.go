package crab

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	_ "image/png"
)

// Constant values like static screen positions are defined here.
const (
	ScreenWidth  = 1000
	ScreenHeight = 800
)

// Game holds our data required for managing state. All data, like images, object positions, and scores, go here.
type Game struct {
}

// NewGame prepares a fresh game state required for startup.
func NewGame() *Game {
	return &Game{}
}

// Update processes all games rules, like checking user input and keeping score. All state updates must occur here, NOT in Draw.
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// Signal that the game shall terminate normally when the Escape key is pressed.
		return ebiten.Termination
	}

	return nil
}

// Draw renders all game images to the screen according to the current game state.
func (g *Game) Draw(screen *ebiten.Image) {
	drawText(screen, 100, 300, color.White, "Project setup succeeded! Press Esc to exit and start programming.")
}

// Layout returns the logical screen size of the game. It can differ from the native outside size and will be scaled if needed.
func (g *Game) Layout(width, height int) (screenWidth, screenHeight int) {
	// No need to use a different logical screen size here. Our game size shall match the native outside window.
	return width, height
}
