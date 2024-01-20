package main

import (
	"fmt"
)

func main() {
	newBlockchain := NewBlockchain()

	newBlockchain.AddBlock("First transaction")
	newBlockchain.AddBlock("Second transaction")
	for i, block := range newBlockchain.Blocks {
		fmt.Printf("Block ID: %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the Block: %x\n", block.MyBlockHash)
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash)
		fmt.Printf("All the transactions: %s\n", block.AllData)
	}
}