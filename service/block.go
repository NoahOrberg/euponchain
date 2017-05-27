package service

import (
	"errors"
	"time"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
)

func getLastInsertedBlock() (model.Block, error) {
	if len(memory.CHAIN)-1 < 0 {
		return model.Block{}, errors.New("No such Block!")
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

func IsValidNewBlock(newBlock model.Block, previousBlock model.Block) bool {
	res = true
	if previousBlock.Index+1 != newBlock.Index {
		res = false
	} else if previousBlock.Hash != newBlock.PreviousHash {
		res = false
	} else if CalcHash(newBlock) != newBlock.Hash {
		res = false
	}
	return res
}

func ReplaceChain(newChain model.Block) {
	if IsValidNewBlock(newChain) && len(newChain) > len(memory.CHAIN) {
		memory.CHAIN = newChain
	}
}
