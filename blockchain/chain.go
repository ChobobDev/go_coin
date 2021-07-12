package blockchain

import (
	"sync"

	"github.com/ChobobDev/go_coin/db"
	"github.com/ChobobDev/go_coin/utils"
)

type blockchain struct {
	NewstHash string `json:"newestHash"`
	Height    int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewstHash, b.Height)
	b.NewstHash = block.Hash
	b.Height = block.Height

}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Jeong Block")
		})
	}
	return b
}
