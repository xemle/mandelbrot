package main

import (
	"encoding/json"
	"mandelbrot/server"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tileCommand server.TileCommand
	err := decoder.Decode(&tileCommand)
	if err != nil {
		panic("Could not parse tileCommand")
	}
	var result = tileCommand.Calculate()
	enc := json.NewEncoder(w)
	enc.Encode(result)
}

func main() {
	http.HandleFunc("/tile", handler)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":8080", nil)
}
