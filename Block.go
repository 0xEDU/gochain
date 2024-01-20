package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Generate hash for a block
// Concatenate all data and hash it
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

// Create a new Block and set it's hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
