package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

type Cryptoblock struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

type BlockChain struct {
	blocks []*Cryptoblock
}

func (c *Cryptoblock) BuildHash() {
	details := bytes.Join([][]byte{
		c.Data,
		c.PrevHash,
		IntToHex(int64(c.Nonce)),
	}, []byte{})
	hash := sha256.Sum256(details)
	c.Hash = hash[:]
}

func ProofOfWork(block *Cryptoblock, difficulty int) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty)) // Adjust the difficulty

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		block.Nonce = nonce
		data := bytes.Join([][]byte{
			block.Data,
			block.PrevHash,
			IntToHex(int64(block.Nonce)),
		}, []byte{})

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			// If hash is less than the target, the proof of work is done
			break
		} else {
			nonce++
		}
	}

	block.Hash = hash[:]
	block.Nonce = nonce
}

func BuildBlock(data string, prevHash []byte, difficulty int) *Cryptoblock {
	block := &Cryptoblock{[]byte{}, []byte(data), prevHash, 0}
	ProofOfWork(block, difficulty)
	return block
}

func (chain *BlockChain) AddBlock(data string, difficulty int) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := BuildBlock(data, prevBlock.Hash, difficulty)
	chain.blocks = append(chain.blocks, new)
}

func Inception(difficulty int) *Cryptoblock {
	return BuildBlock("Inception", []byte{}, difficulty)
}

func InitializeBlockChain(difficulty int) *BlockChain {
	return &BlockChain{[]*Cryptoblock{Inception(difficulty)}}
}

func IntToHex(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		fmt.Println("Error converting int to hex:", err)
	}
	return buff.Bytes()
}

func main() {
	difficulty := 20 // Adjust the difficulty as needed Controls how hard it is to find a valid hash. Higher difficulty means more work required.
	chain := InitializeBlockChain(difficulty)

	for i := 1; i <= 10; i++ {
		chain.AddBlock(fmt.Sprintf("%d block", i), difficulty)
	}

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
