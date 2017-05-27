package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
	"github.com/NoahOrberg/euponchain/service"
)

func BlocksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBlocksHandler(w, r)
	case http.MethodPost:
		AddBlocksHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AddBlocksHandler(w http.ResponseWriter, r *http.Request) {
	// JSON で受け取る（文字列Body）
	data := model.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// ハッシュ化して、ブロックを作る
	block, err := service.CreateBlock(data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// ネットワークへ送信
	err = service.SendNewBlock(block)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(memory.CHAIN); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
