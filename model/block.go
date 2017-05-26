package model

type Block struct {
	Index        int64
	PreviousHash []byte
	Timestamp    int64
	Data         Data
	Hash         []byte
}

type Data struct {
	Body string `json:"body"`
}
