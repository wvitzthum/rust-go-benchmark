package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	graph := exampleGraph()
	res := graph.BFS(5)
	w.Write([]byte(fmt.Sprintf("u: %d, v: %d", res[0], res[1])))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/graph", graphHandler)
	http.ListenAndServe(":8080", nil)
}