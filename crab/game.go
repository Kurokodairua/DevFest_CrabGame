package crab

import (
	"image"
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/isensee-bastian/crab/resources/images/sprites"
)

// Constant values like static screen positions are defined here.
const (
	ScreenWidth  = 1000
	ScreenHeight = 800
)

// Game holds our data required for managing state. All data, like images, object positions, and scores, go here.
type Game struct {
	beach *ebiten.Image
	crab  []*ebiten.Image // Array für Crab Frames
	fish  *ebiten.Image
	crabX int
	crabY int
	fishX int
	fishY int
}

// NewGame prepares a fresh game state required for startup.
func NewGame() *Game {
	return &Game{
		beach: readImage(sprites.Beach),
		fish:  readImage(sprites.Fish),
		crab:  readAnimationImages(sprites.Crab),
		crabX: 400,
		crabY: 500,
		fishX: 300,
		fishY: 400,
	}
}

// Update processes all games rules, like checking user input and keeping score. All state updates must occur here, NOT in Draw.
func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		// Signal that the game shall terminate normally when the Escape key is pressed.
		return ebiten.Termination
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
		g.crabX = min(g.crabX+3, ScreenWidth-spriteWidth-1) // ohne spriteWidth kann die Krabbe shoulder peaken
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
		g.crabX = max(g.crabX-3, 0)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		g.crabY = max(g.crabY-1, 0)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
		g.crabY = min(g.crabY+1, ScreenHeight-spriteHeight-1)
	}
	crabRect := image.Rect(g.crabX, g.crabY, g.crabX+spriteWidth, g.crabY+spriteHeight)
	fishRect := image.Rect(g.fishX, g.fishY, g.fishX+spriteWidth, g.fishY+spriteHeight)
	if crabRect.Overlaps(fishRect) {
		g.fishX += rand.Intn(ScreenWidth+1-spriteWidth) + spriteWidth
		g.fishY += rand.Intn(ScreenHeight+1-spriteHeight) + spriteHeight //TODO Fisch despawnt :(
	}

	return nil
}

// Draw renders all game images to the screen according to the current game state.
func (g *Game) Draw(screen *ebiten.Image) {
	// Als erstes den Hintergrund:
	{
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(0, 0) //wo soll das Bild dargestellt werden
		opts.GeoM.Scale(2, 2)     //um wieviel wollen wir das Bild vergrößern
		screen.DrawImage(g.beach, opts)
	}
	// Alle anderen Sprites werden darüber gelagert:
	{
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(g.crabX), float64(g.crabY)) //wo soll das Bild dargestellt werden
		screen.DrawImage(g.crab[0], opts)                       //TODO: nicht nur erster Frame
	}
	{
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(g.fishX), float64(g.fishY))
		screen.DrawImage(g.fish, opts)
	}
}

// Layout returns the logical screen size of the game. It can differ from the native outside size and will be scaled if needed.
func (g *Game) Layout(width, height int) (screenWidth, screenHeight int) {
	// No need to use a different logical screen size here. Our game size shall match the native outside window.
	return width, height
}
