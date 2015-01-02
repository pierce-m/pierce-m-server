package main

import (
	"fmt"
)

func ring_handler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/ring", handler)
	http.ListenAndServe(":9090", nil)
}
