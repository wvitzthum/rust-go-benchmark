package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

type handler struct{}
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type subHandler struct {
	val atomic.Value
}
func (h *subHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.val.Load().(bool) {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

var (
	ready = atomic.Value{}
)

func main() {
	ready.Store(false)

	mux := http.NewServeMux()
	mux.Handle("/bench", &handler{})
	mux.Handle("/readiness", &subHandler{val: ready})
	go http.ListenAndServe(":8080", mux)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	ready.Store(true)
	<-sigChan
	log.Printf("Terminating service")

}