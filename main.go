package main

import "go-beep-shoutcast/stream"

func main() {
	done := make(chan bool)
	teardown, err := stream.Play("http://stream.radioparadise.com/mp3-192")
	if err != nil {
		panic(err)
	}
	defer teardown()

	// Wait audio to finish playing
	<-done
}
