package main

import (
	"go.etcd.io/bbolt"
	"log"
)

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	db          *bbolt.DB
	currentHash []byte
}

// Next returns the next block in the blockchain
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		serialized := b.Get(i.currentHash)
		if serialized == nil {
			return nil
		}
		block = DeserializeBlock(serialized)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	if block != nil {
		i.currentHash = block.PrevBlockHash
	}

	return block
}

// HasNext checks if there are more blocks to iterate
func (i *BlockchainIterator) HasNext() bool {
	return len(i.currentHash) > 0
}