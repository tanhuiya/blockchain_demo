package main

import (
	"fmt"
)

func (cli *CLI) ReindexUTXO() {
	bc := NewBlockChain()
	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transanctions in the UTXO set .\n", count)
}
