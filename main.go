package main

import (
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/dolmen-go/kittyimg"

	"go-beep-shoutcast/shoutcast"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	stream, err := shoutcast.Open("http://stream.radioparadise.com/mp3-128")
	if err != nil {
		panic(err)
	}

	// Callback function to be called when song changes
	stream.MetadataCallbackFunc = func(m *shoutcast.Metadata) {
		println("Now listening to: ", m.StreamTitle)
		println("StreamURL: ", m.StreamURL)
		if m.StreamURL != "" {
			resp, err := http.Get(m.StreamURL)
			if err != nil {
				println("[ERROR] could not get ", m.StreamURL)
				return
			}
			img, _, err := image.Decode(resp.Body)
			if err != nil {
				println("[ERROR] could not decode image")
				return
			}
			kittyimg.Fprintln(os.Stdout, img)

		}
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
