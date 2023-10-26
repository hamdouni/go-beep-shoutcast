package main

import (
	"net/http"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	// stream, err := shoutcast.Open("http://stream.radioparadise.com/mp3-128")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Callback function to be called when song changes
	// stream.MetadataCallbackFunc = func(m *shoutcast.Metadata) {
	// 	println("Now listening to: ", m.StreamTitle)
	// }
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
	err = speaker.Init(format.SampleRate, 8192)
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
