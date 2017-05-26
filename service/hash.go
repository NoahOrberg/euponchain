package service

import (
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
)

func CalcHash(data model.Block) []byte {
	byteID := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(byteID, data.Index)
	byteTimestamp := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(byteTimestamp, data.Timestamp)

	hdata := []byte{}
	hdata = append(hdata, byteID...)
	hdata = append(hdata, byteTimestamp...)
	hdata = append(hdata, data.PreviousHash...)
	res := sha256.Sum256(hdata)
	return res[:]
}

func getLastInsertedBlock() (model.Block, error) {
	if len(memory.CHAIN)-1 < 0 {
		res := model.Block{
			Index: 0,
			Hash:  []byte("Nil"),
		}
		return res, nil
	}
	return memory.CHAIN[len(memory.CHAIN)-1], nil
}

func CreateBlock(data model.Data) (model.Block, error) {
	previousBlock, err := getLastInsertedBlock()
	if err != nil {
		return model.Block{}, nil
	}
	newBlock := model.Block{
		Index:        previousBlock.Index + 1,
		PreviousHash: previousBlock.Hash,
		Timestamp:    time.Now().Unix(),
		Data:         data,
	}
	newBlock.Hash = CalcHash(newBlock)
	return newBlock, nil
}
