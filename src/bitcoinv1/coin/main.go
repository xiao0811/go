package main

import (
	"bitcoinv1/core"
)

func main() {
	bc := core.NewBlockChain()
	defer bc.Db.Close()

	cli := core.CLI{Bc:bc}
	cli.Run()
}
