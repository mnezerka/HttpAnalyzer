package main

import (
    "io/ioutil"
    "net/http"
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

    // Loop through headers
    h.log.Debugf("Headers:")
    for name, headers := range r.Header {
        for _, val := range headers {
            h.log.Debugf("  %v: %v", name, val)
        }
    }

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        h.log.Errorf("Reading request body error: %s", err)
    }
    h.log.Debugf("Read request body passed, size: %d bytes", len(body))
}
