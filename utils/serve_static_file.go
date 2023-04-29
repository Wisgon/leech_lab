package utils

import (
	"log"
	"net/http"
)

func ServerStaticFile(folderPath string) {
	fs := http.FileServer(http.Dir(folderPath))
	http.Handle("/", fs)

	log.Print("Listening on :8002...")
	err := http.ListenAndServe(":8002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
