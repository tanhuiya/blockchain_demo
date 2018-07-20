package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) PrintChain() {
	bc := NewBlockChain()
	defer bc.db.Close()
	bci := bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}

		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
