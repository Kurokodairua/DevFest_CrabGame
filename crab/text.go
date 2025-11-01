package crab

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"log"
)

var (
	normalTextFace *text.GoTextFace
	bigTextFace    *text.GoTextFace
)

// init sets up text prerequisites automatically. It must not be called directly.
func init() {
	source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatalf("Error on creating new text face source: %v", err)
	}

	normalTextFace = &text.GoTextFace{
		Source: source,
		Size:   24,
	}
	bigTextFace = &text.GoTextFace{
		Source: source,
		Size:   36,
	}
}

// drawText is a helper function for printing a normal-sized text string on the screen.
func drawText(screen *ebiten.Image, x, y int, color color.Color, content string) {
	drawTextWithFace(screen, x, y, color, content, normalTextFace)
}

// drawBigText is a helper function for printing a large-sized text string on the screen.
func drawBigText(screen *ebiten.Image, x, y int, color color.Color, content string) {
	drawTextWithFace(screen, x, y, color, content, bigTextFace)
}

// drawTextWithFace can be used for printing a more custom text string on the screen. When possible, using drawText or
// drawBigText should be preferred as it is simpler to use.
func drawTextWithFace(screen *ebiten.Image, x, y int, color color.Color, content string, textFace *text.GoTextFace) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color)
	text.Draw(screen, content, textFace, op)
}
