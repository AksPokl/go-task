package handler

import (
	"log"
	"net/http"
)

func HandleError(err error, w http.ResponseWriter, status int) {
	log.Printf("HTTP %d - %s", status, err.Error)
	http.Error(w, err.Error(), status)
	w.WriteHeader(status)
}
