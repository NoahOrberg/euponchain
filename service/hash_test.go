package service

import (
	"fmt"
	"testing"

	"github.com/NoahOrberg/euponchain/model"
)

func TestCalcHash(t *testing.T) {
	data := model.Block{
		Index:        12,
		PreviousHash: []byte("AAA"),
		Timestamp:    12,
		Data: model.Data{
			Body: "FIRST !",
		},
	}
	hash := CalcHash(data)

	fmt.Print("Hash: %v", string(hash))
}
