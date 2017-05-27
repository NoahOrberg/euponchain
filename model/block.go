package model

type Block struct {
	Index        int64  `json:"index"`
	PreviousHash []byte `json:"previousHash"`
	Timestamp    int64  `json:"timestamp"`
	Data         Data   `json:"data"`
	Hash         []byte `json:"hash"`
}

type Data struct {
	Body string `json:"body"`
}
