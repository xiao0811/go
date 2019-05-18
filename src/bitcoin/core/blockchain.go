package core

// BlockChain keeps a sequence of Blocks
type BlockChain struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the blockChain
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
