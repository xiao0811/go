package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64  // 区块创建时间戳
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 前一个区块的哈希值
	Hash          []byte // 区块自身的哈希值, 用于校验区块数据有效
	Nonce         int
}

// NewBlock creates and returns Block
// 初始化创建一个新区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// SetHash calculates and sets block hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}

	return &block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}
