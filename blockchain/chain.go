package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/ChobobDev/go_coin/db"
	"github.com/ChobobDev/go_coin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.HandleErr(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			fmt.Printf("NH : %s \n H : %d \n", b.NewestHash, b.Height)
			// db로 부터 checkpoint를 찾은후 - byte로부터 b를 복구
			persistedBlockchain := db.Blockchain()
			if persistedBlockchain == nil {
				b.AddBlock("Bernie the great Gopher block")
			} else {
				fmt.Println("Restoring taking in place")
				b.restore(persistedBlockchain)
			}
			fmt.Printf("NH : %s \n H : %d \n", b.NewestHash, b.Height)

		})
	}
	return b
}
