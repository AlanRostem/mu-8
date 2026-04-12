package multimedia

import (
	"math"

	"github.com/AlanRostem/mu-8/internal/logger"
	"github.com/AlanRostem/mu-8/interpreter"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const (
	sampleRate = 48000
	frequency  = 440
)

// stream is an infinite stream of 440 Hz sine wave.
// Source: https://ebitengine.org/en/examples/sinewave.html#Code
type stream struct {
	pos int64
}

// Read is io.Reader's Read.
//
// Read fills the data with sine wave samples.
func (s *stream) Read(buf []byte) (int, error) {
	const bytesPerSample = 8

	n := len(buf) / bytesPerSample * bytesPerSample

	const length = sampleRate / frequency
	for i := 0; i < n/bytesPerSample; i++ {
		v := math.Float32bits(float32(math.Sin(2 * math.Pi * float64(s.pos/bytesPerSample+int64(i)) / length)))
		buf[8*i] = byte(v)
		buf[8*i+1] = byte(v >> 8)
		buf[8*i+2] = byte(v >> 16)
		buf[8*i+3] = byte(v >> 24)
		buf[8*i+4] = byte(v)
		buf[8*i+5] = byte(v >> 8)
		buf[8*i+6] = byte(v >> 16)
		buf[8*i+7] = byte(v >> 24)
	}

	s.pos += int64(n)
	s.pos %= length * bytesPerSample

	return n, nil
}

// Close is io.Closer's Close.
func (s *stream) Close() error {
	return nil
}

type SoundEngine struct {
	interpreter  *interpreter.Interpreter
	audioContext *audio.Context
	player       *audio.Player

	isPlaying bool
}

func NewSoundEngine(interpreter *interpreter.Interpreter) *SoundEngine {
	audioContext := audio.NewContext(sampleRate)
	return &SoundEngine{
		interpreter:  interpreter,
		audioContext: audioContext,
		player:       nil,
	}
}

func (s *SoundEngine) Update() {
	if s.interpreter.SoundTimer() > 0 {
		if !s.isPlaying {
			s.isPlaying = true
			player, err := s.audioContext.NewPlayerF32(&stream{})
			if err != nil {
				panic(err)
			}
			s.player = player
			s.player.Play()
		}
	} else if s.isPlaying {
		s.isPlaying = false
		err := s.player.Close()
		if err != nil {
			logger.Errorf("%v", err)
		}
	}
}
