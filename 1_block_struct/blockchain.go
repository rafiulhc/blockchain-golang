package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce int
	previousHash string
	timeStamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timeStamp = time.Now().UnixNano()
	return b
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("timeStamp: %d\n", b.timeStamp)
	fmt.Printf("transactions: %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "")
	b.Print()
}