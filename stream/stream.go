package stream

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

var (
	stream   io.ReadCloser
	streamer beep.StreamSeekCloser
)

func Play(URL string) (TearDown func(), err error) {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	stream = resp.Body

	// Decode audio
	var format beep.Format
	streamer, format, err = mp3.Decode(stream)
	if err != nil {
		return TearDown, fmt.Errorf("stream: decoding mp3 %s", err)
	}

	// Initialize speaker
	err = speaker.Init(format.SampleRate, 16384)
	if err != nil {
		return TearDown, fmt.Errorf("stream: initializing speaker %s", err)
	}

	// Play audio
	speaker.Play(streamer)

	return TearDown, nil
}

func TearDown() {
	stream.Close()
	streamer.Close()
}
