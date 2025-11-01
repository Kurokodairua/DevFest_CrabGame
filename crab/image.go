package crab

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const (
	spriteWidth           = 192 / 4
	spriteHeight          = 192 / 4
	animationFrameColumns = 4
)

// readImage takes raw image bytes and returns a ready-to-use image for drawing. It is expected that the passed image
// bytes contain a single image only.
func readImage(rawImage []byte) *ebiten.Image {
	stdImage, _, err := image.Decode(bytes.NewReader(rawImage))
	if err != nil {
		log.Fatalf("Error while loading image: %v", err)
	}
	return ebiten.NewImageFromImage(stdImage)
}

// readAnimationImages takes raw image bytes that represent multiple images in a single row. It is intended for
// animations and therefore returns all images from the specified row, ready-to-use for drawing individually.
func readAnimationImages(rawImage []byte) []*ebiten.Image {
	stdAnimationImage, _, err := image.Decode(bytes.NewReader(rawImage))
	if err != nil {
		log.Fatalf("Error while loading image: %v", err)
	}
	animationImage := ebiten.NewImageFromImage(stdAnimationImage)

	var allFrames []*ebiten.Image

	for index := 0; index < animationFrameColumns; index++ {
		xOffset := index * spriteWidth
		frameImage := animationImage.SubImage(image.Rect(
			xOffset,
			0,
			xOffset+spriteWidth-1,
			spriteHeight-1,
		))
		allFrames = append(allFrames, ebiten.NewImageFromImage(frameImage))
	}

	return allFrames
}
