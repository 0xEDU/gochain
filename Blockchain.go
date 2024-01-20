package main

// Add a new block to Blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// Create a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
