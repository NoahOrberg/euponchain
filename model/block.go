package model

type Block struct {
	Index        int64
	previousHash []byte
	Timestamp    int64
	data         Data
	Hash         []byte
}

type Data struct {
	body string
}
