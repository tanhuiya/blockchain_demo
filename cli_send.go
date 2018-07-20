package main

import (
	"fmt"
	"log"
)

func (cli *CLI) Send(from, to string, amount int) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}
	bc := NewBlockChain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbtx := NewCoinbaseTX(from, "")
	block := bc.MineBlock([]*Transaction{cbtx, tx})
	UTXOSet.Update(block)
	fmt.Println("success !")
}
