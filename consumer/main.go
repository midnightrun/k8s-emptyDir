package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func handleListen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		select {
		case <-r.Context().Done():
			return
		default:
			file, err := ioutil.ReadFile("/static/exchange")
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "error reading from file: %s\n", err)
				file = []byte(err.Error())
			}
			w.Write(formatSSE(
				"message",
				string(file)))
			w.(http.Flusher).Flush()
		}

		time.Sleep(time.Second)
	}
}

func formatSSE(event, message string) []byte {
	eventPayload := "event: " + event + "\n"

	dataLines := strings.Split(message, "\n")

	for _, line := range dataLines {
		eventPayload = eventPayload + "data: " + line + "\n"
	}

	return []byte(eventPayload + "\n")
}

func main() {
	exePath, err := os.Executable()
	exeDir := filepath.Dir(exePath)
	fs := http.FileServer(http.Dir(filepath.Join(exeDir, "/assets")))
	http.Handle("/", fs)
	http.HandleFunc("/listen", handleListen)

	log.Println("Listening on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
