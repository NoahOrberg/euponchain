package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/NoahOrberg/euponchain/config"
	"github.com/NoahOrberg/euponchain/controller"
	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var config config.Config
	err := envconfig.Process("eupon", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Nodes TODO:自動で追加にする
	memory.Nodes = []model.Node{
		{
			Host: "localhost",
			Port: 9090,
		},
		{
			Host: "localhost",
			Port: 9091,
		},
		{
			Host: "localhost",
			Port: 9092,
		},
		{
			Host: "localhost",
			Port: 9093,
		},
	}
	//

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello:)")
	})
	r.HandleFunc("/blocks", controller.BlocksHandler)
	r.HandleFunc("/blocks/nodes", controller.NodesHandler)

	port := ":" + strconv.FormatInt(config.Port, 10)
	log.Println("http://localhost" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
