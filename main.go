package main

import (
	"bytes"
	"fmt"

	"crypto/sha256"
)

type Cryptoblock struct {
	Hash []byte

	Data []byte

	PrevHash []byte
}
type BlockChain struct {
	blocks []*Cryptoblock
}

func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][]byte{c.Data, c.PrevHash}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func BuildBlock(data string, prevHash []byte) *Cryptoblock {
	block := &Cryptoblock{[]byte{}, []byte(data), prevHash}
	block.BuildHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := BuildBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Inception() *Cryptoblock {
	return BuildBlock("Inception", []byte{})
}

func InitializeBlockChain() *BlockChain {
	return &BlockChain{[]*Cryptoblock{Inception()}}
}
func main() {
	chain := InitializeBlockChain()
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, c := range l {
		chain.AddBlock(fmt.Sprintf("%d block", c))
	}

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
