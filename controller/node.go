package controller

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusBadRequest)
	}
	if service.IsValidNewBlock(block, pBlock) {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// ブロックをmemory.CHAIN へインサートする
		memory.CHAIN = append(memory.CHAIN, block)
	}
}
