package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func ring_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ioutil.WriteFile("infile", []byte(r.FormValue("program")), 0644)
	resp, err := exec.Command("ring", "infile").Output()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		fmt.Fprintf(w, "%s", string(resp))
	}
}

func main() {
	http.HandleFunc("/ring/", ring_handler)
	http.ListenAndServe(":9090", nil)
}
