package service

import (
	"crypto/sha256"
	"encoding/binary"

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
