package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
	"github.com/NoahOrberg/euponchain/service"
)

func NodesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddBlocksNodesHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func AddBlocksNodesHandler(w http.ResponseWriter, r *http.Request) {
	// JSON で受け取る（文字列Body）
	block := model.Block{}
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// 正当か判断
	// 不当ならレスポンスで不当な旨伝える
	pBlock, err := service.GetLastInsertedBlock()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	log.Printf("%v\n", service.IsValidNewBlock(block, pBlock))
	if service.IsValidNewBlock(block, pBlock) {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// ブロックをmemory.CHAIN へインサートする
		memory.CHAIN = append(memory.CHAIN, block)
		log.Println("Inserted ")
	}
}
