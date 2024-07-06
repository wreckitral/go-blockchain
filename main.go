package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
    Timestamp int64 // Timestamp is the timestamp when the block is created
    Data []byte // Data is the actual information
    PrevBlockHash []byte // PrevBlockHash is the hash of previous block
    Hash []byte // Hash is the hash of the block
}

// the original way of blockchain to calculate their hash is computationally difficult operation thats why blockchain is secure, it takes some time even with fast computers.
// to make things simple, i just take block fields, concatenate them, and then calculate a SHA-256 on them.

func (b *Block) SetHash() {
    timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

    headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

    hash := sha256.Sum256(headers)

    b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
    block := &Block{
        Timestamp: time.Now().Unix(),
        Data: []byte(data),
        PrevBlockHash: prevBlockHash,
        Hash: []byte{},
    }

    block.SetHash()

    return block
}

type Blockchain struct {
    blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks)-1]

    newBlockk := NewBlock(data, prevBlock.Hash)

    bc.blocks = append(bc.blocks, newBlockk)
}


func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}







