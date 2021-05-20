package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)

// for preview usage, delete if development progress is started

func httpdemo1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World, "+r.RemoteAddr)
}
func main() {
	http.HandleFunc("/", httpdemo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
