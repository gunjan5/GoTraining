package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", MeowServer)
	http.ListenAndServe(":8080", nil)

}

func MeowServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Meow, %s!", r.URL.Path[1:])
}
