package main

import (
	"log"
	"net/http"

	"go-beep-shoutcast/stream"
)

var (
	err   error
	count = 0
	audio struct {
		teardown func()
	}
)

func main() {
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)
	http.ListenAndServe(":7334", nil)
}

func start(w http.ResponseWriter, r *http.Request) {
	count++
	audio.teardown, err = stream.Play("http://stream.radioparadise.com/mp3-192")
	if err != nil {
		panic(err)
	}
	log.Printf("teardown: %#v\n", count)
}

func stop(w http.ResponseWriter, r *http.Request) {
	if audio.teardown != nil {
		audio.teardown()
		return
	}
	log.Printf("teardown: %#v\n", count)
}
