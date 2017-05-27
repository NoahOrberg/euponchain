package memory

import "github.com/NoahOrberg/euponchain/model"

var first = model.Block{
	Index:        0,
	PreviousHash: []byte("0"),
	Timestamp:    0,
	Data: model.Data{
		Body: "GENESIS BLOCK",
	},
	Hash: []byte("816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7"),
}
var CHAIN = []model.Block{first}

var Nodes = []model.Node{}
