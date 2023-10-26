package main

import (
	"net/http"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func main() {
	streamURL := "http://stream.radioparadise.com/mp3-192"
	resp, err := http.Get(streamURL)
	if err != nil {
		panic(err)
	}
	stream := resp.Body
	defer stream.Close()

	// Decode audio
	streamer, format, err := mp3.Decode(stream)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = streamer.Close(); err != nil {
			panic(err)
		}
	}()

	// Initialize speaker
	err = speaker.Init(format.SampleRate, 16384)
	if err != nil {
		panic(err)
	}

	// Play audio
	println("playing ", streamURL)
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait audio to finish playing
	<-done
}
