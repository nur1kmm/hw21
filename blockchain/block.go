package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates and returns a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	
	// Mine the block with proof of work
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

// NewGenesisBlock creates and returns the genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// SetHash calculates and sets the hash of the block (deprecated with PoW)
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}