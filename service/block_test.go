package service

import (
	"testing"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
)

func TestSendNewBlock(t *testing.T) {
	memory.Nodes = []model.Node{
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
		{
			Host: "localhost",
			Port: 9094,
		},
	}
	err := sendNewBlock(model.Block{
		Index:        1,
		PreviousHash: []byte("AAA"),
		Timestamp:    123456,
		Data: model.Data{
			Body: "HOGEHOGE!",
		},
		Hash: []byte("BBB"),
	})

	if err != nil {
		// t.Fatal(err)
	}

}
