package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

const (
	defaultPort = ":8181"
	logInFile   = true
	logFilename = "events.log"
)

var giphyAPIKey string

func init() {
	envKey, ok := os.LookupEnv("GIPHY_APIKEY")
	if !ok {
		log.Fatalln("could not retrieve GIPHY_APIKEY from ENV")
	}
	giphyAPIKey = envKey
}

func main() {
	r := createRouter()
	w := createLog()
	logR := handlers.LoggingHandler(w, r)

	err := http.ListenAndServe(defaultPort, logR)
	if err != nil {
		log.Fatal(err)
	}
}

func createLog() io.Writer {
	w := os.Stderr
	flags := os.O_APPEND | os.O_WRONLY

	if logInFile {
		if _, exists := os.Stat(logFilename); exists == nil {
			w, _ = os.OpenFile(logFilename, flags, 0644)
		} else if os.IsNotExist(exists) {
			w, _ = os.OpenFile(logFilename, os.O_CREATE|flags, 0644)
		} else {
			log.Fatalln("could not create log file:", exists.Error())
		}
	}
	return w
}
