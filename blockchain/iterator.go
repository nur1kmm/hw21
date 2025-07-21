package blockchain

import (
	"go.etcd.io/bbolt"
)

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bbolt.DB
}

// Next returns the next block from the blockchain
func (i *BlockchainIterator) Next() (*Block, error) {
	var block *Block

	err := i.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		if encodedBlock == nil {
			return ErrBlockNotFound
		}

		block = DeserializeBlock(encodedBlock)
		i.currentHash = block.PrevBlockHash

		return nil
	})

	if err != nil {
		return nil, err
	}

	return block, nil
}
