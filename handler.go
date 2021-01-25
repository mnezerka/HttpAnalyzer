package main

import (
    "net/http"
    "io/ioutil"
    "github.com/op/go-logging"
)

type Handler struct {
    log *logging.Logger
}

func NewHandler(log *logging.Logger) *Handler {
    return &Handler{log: log}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    h.log.Debugf("Incoming packet uri: %s, method: %s, content length: %d", r.RequestURI, r.Method, r.ContentLength)

	if r.Method == "POST" {
		bodyText, err := ioutil.ReadAll(r.Body)
		if err != nil {
			h.log.Fatal("Cannot decode request body")
		}
		h.log.Debugf("Body: %s", string(bodyText))
	}
}
