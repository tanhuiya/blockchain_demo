package main

import "fmt"

func (cli *CLI) createBlockChain(address string) {
	bc := CreateBlockChain(address)
	bc.db.Close()
	fmt.Println("Done !")
}
