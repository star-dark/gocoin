package main

import (
	"fmt"

	"github.com/star-dark/gocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("second Block")
	chain.AddBlock("third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlock() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
