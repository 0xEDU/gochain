package main

import (
	"fmt"
)

func main() {
	newBlockchain := NewBlockchain()

	newBlockchain.AddBlock("First transaction")
	newBlockchain.AddBlock("Second transaction")
	for _, block := range newBlockchain.Blocks {
		fmt.Printf("Hash of the Block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %s\n", block.AllData)
	}
}