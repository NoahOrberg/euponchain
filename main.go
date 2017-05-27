package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NoahOrberg/euponchain/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello:)")
	})
	r.HandleFunc("/blocks", controller.BlocksHandler)
	port := ":8282"
	fmt.Print("http://localhost", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
