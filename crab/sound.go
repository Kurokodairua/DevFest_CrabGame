package crab

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
)

const (
	sampleRate = 48000
)

var audioContext *audio.Context

// init sets up the audio context automatically. It must not be called directly. This is required for audio player creation.
func init() {
	audioContext = audio.NewContext(sampleRate)
}

// AudioPlayer is a thin wrapper around Ebitengines original audio.Player type to simplify some commonly used actions
// like rewinding and playing a sound or closing the sound stream during application shutdown.
type AudioPlayer struct {
	*audio.Player
}

// newMp3AudioPlayer takes a raw mp3 file as bytes and returns a ready-to-use audio player. Note that Close() MUST be
// called for proper cleanup when the audio player is no longer needed, typically at application shutdown.
func newMp3AudioPlayer(rawSound []byte) *AudioPlayer {
	sellSound, err := mp3.DecodeF32(bytes.NewReader(rawSound))
	if err != nil {
		log.Fatalf("Failed to decode raw sound as mp3: %v", err)
	}

	audioPlayer, err := audioContext.NewPlayerF32(sellSound)
	if err != nil {
		log.Fatalf("Failed to create mp3 audio player: %v", err)
	}

	return &AudioPlayer{audioPlayer}
}

// Replay resets the played audio to the start and plays it. This makes it easy to re-play a sound as often as needed
// in a single step.
func (a *AudioPlayer) Replay() {
	err := a.Rewind()

	if err != nil {
		// Logging is sufficient here as playing sounds is not critical for the overall gameplay.
		log.Printf("Error on rewinding audio: %v", err)
		return
	}

	a.Play()
}

// Close MUST be called when the audio player is no longer needed, typically when the game exits due to normal termination
// or due to an error, i.e. during game loop cleanup.
func (a *AudioPlayer) Close() {
	if a == nil {
		// Nothing to do.
		return
	}

	err := a.Player.Close()

	// Logging is sufficient here. Keep it simple and avoid unnecessary error handling on the caller side.
	if err != nil {
		log.Printf("Error on closing audio: %v", err)
	}
}
