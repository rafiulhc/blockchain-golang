package main

import (
	"log"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
}

func (b *Block) Print(){
	println("Nonce:", b.nonce)
	println("Previous hash:", b.previousHash)
	println("Timestamp:", b.timestamp)
}

func init (){
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
