package main

import (
	"blockchain"
	"fmt"
)

func main() {
	blockChain := blockchain.InitBlockChain()

	blockChain.AddBlock("First block")
	blockChain.AddBlock("Second block")
	blockChain.AddBlock("Third block")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Previuos Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
