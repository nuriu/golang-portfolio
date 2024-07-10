package main

import (
	"fmt"
	"net/http"
	"rest-api/configs"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello /")
	})

	addr := configs.Environment.Host + ":" + configs.Environment.Port
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err.Error())
	}
}
