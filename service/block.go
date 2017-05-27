package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/NoahOrberg/euponchain/memory"
	"github.com/NoahOrberg/euponchain/model"
)

func GetLastInsertedBlock() (model.Block, error) {
	if len(memory.CHAIN)-1 < 0 {
		return model.Block{}, errors.New("No such Block!")
	}
	return memory.CHAIN[len(memory.CHAIN)-1], nil
}

func CreateBlock(data model.Data) (model.Block, error) {
	previousBlock, err := GetLastInsertedBlock()
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
	res := true
	if previousBlock.Index+1 != newBlock.Index {
		res = false
	} else if reflect.DeepEqual(previousBlock.Hash, newBlock.PreviousHash) {
		res = false
	} else if reflect.DeepEqual(CalcHash(newBlock), newBlock.Hash) {
		res = false
	}
	return res
}

func IsValidChain(chain []model.Block) bool {
	return true // TODO: あとで実装 (とりあえずTRUEのみ返す)
}

func ReplaceChain(newChain []model.Block) {
	if IsValidChain(newChain) && len(newChain) > len(memory.CHAIN) {
		memory.CHAIN = newChain
	}
}

func SendNewBlock(newBlock model.Block) error {
	for _, node := range memory.Nodes {
		nodeUrl := "http://" + node.ToString() + "/blocks/nodes/"
		jsonNewBlock, err := json.Marshal(newBlock)
		if err != nil {
			return err
		}
		log.Print(nodeUrl + "\n")
		log.Print(string(jsonNewBlock) + "\n")
		req, err := http.NewRequest(
			"POST",
			nodeUrl,
			bytes.NewBuffer(jsonNewBlock),
		)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return errors.New("cannot OK")
		}
		res.Body.Close()
	}
	return nil
}
