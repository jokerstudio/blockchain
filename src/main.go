package main

import (
	"blockchain"
	"fmt"
	"strconv"
)

func main() {
	blockChain := blockchain.InitBlockChain()

	blockChain.AddBlock("First block")
	blockChain.AddBlock("Second block")
	blockChain.AddBlock("Third block")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Previuos Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
