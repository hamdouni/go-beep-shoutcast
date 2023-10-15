package main

import (
	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/romantomjak/shoutcast"
)

func main() {
	stream, err := shoutcast.Open("http://stream.radioparadise.com/mp3-128")
	if err != nil {
		panic(err)
	}

	// Callback function to be called when song changes
	stream.MetadataCallbackFunc = func(m *shoutcast.Metadata) {
		println("Now listening to: ", m.StreamTitle)
	}

	// Decode audio
	streamer, format, err := mp3.Decode(stream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize speaker
	err = speaker.Init(format.SampleRate, 8192)
	if err != nil {
		log.Fatal(err)
	}

	// Play audio
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Wait audio to finish playing
	<-done
}
