package controller

import (
	"encoding/json"
	"net/http"

	"github.com/NoahOrberg/euponchain/model"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	// JSON で受け取る（文字列Body）
	data := model.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// ハッシュ化して、ブロックを作る

	// ブロックをインサートする
}
