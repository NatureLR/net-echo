package netecho

import (
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}
}
