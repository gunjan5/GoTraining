package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Pair struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

var storage map[string]string

func handleOperation(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var p Pair
	err := decoder.Decode(&p)

	if err != nil {
		fmt.Printf("[Error] bah: %s", err)
		os.Exit(1)
	}

	switch request.Method {
	case "GET":
		_, ok := storage[p.Key]
		if !ok {
			fmt.Fprintf(w, "Key (%s) doesn't exist in the datastore\n", p.Key)
		}
		fmt.Fprintf(w, "GET: %s = %s\n", p.Key, storage[p.Key])
	case "POST":
		storage[p.Key] = p.Value
		fmt.Fprintf(w, "Successfully POSTed: %s = %s\n", p.Key, p.Value)
	case "DELETE":
		_, ok := storage[p.Key]
		if !ok {
			fmt.Fprintf(w, "Key (%s) doesn't exist in the datastore\n", p.Key)
		}
		fmt.Fprintf(w, "DELETEd: %s = %s\n", p.Key, storage[p.Key])
		delete(storage, p.Key)
	default:
		fmt.Fprintln(w, "invalid operation")
		os.Exit(1)
	}
}

func main() {
	storage = map[string]string{}
	http.HandleFunc("/", handleOperation)
	http.ListenAndServe(":8080", nil)
}
