package main

import (
	"bitcoin/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain() // 初始化区块链, 创建第一个区块(创世纪区块)

	bc.AddBlock("Send 1 BTC to Ivan")      // 加入一个区块, 发送一个比特币给Ivan
	bc.AddBlock("Send 2 more BTC to Ivan") // 加入一个区块, 发送更多比特币给Ivan

	for _, block := range bc.Blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
