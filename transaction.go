package main

import (
	"encoding/hex"
	"fmt"
	"crypto/sha256"
	"log"
	"encoding/gob"
	"bytes"
)

const subsidy = 10

type Transaction struct {
	ID		[]byte
	Vin 	[]TXInput
	Vout	[]TXOutput
}


type TXInput struct {
	// 一个交易输入引用了之前一笔交易的输出，txid 表示引用了之前哪笔交易
	Txid	[]byte
	// 引用之前交易输出的 index
	Vout	int
	// 用来解锁 vout 的数据
	ScriptSig	string
}

type TXOutput struct {
	Value 	int
	ScriptPubKey string
}

func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && 
	len(tx.Vin[0].Txid) == 0 &&
	 tx.Vin[0].Vout == -1
}

func (tx *Transaction) SetID() {
	var encode bytes.Buffer
	var hash [32]byte


	enc := gob.NewEncoder(&encode)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encode.Bytes())
	tx.ID = hash[:]
}

func (in *TXInput) CanUnlockOutPutWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewCoinbaseTX(to, data string) * Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	// 挖矿的奖励， 对应交易为空， 交易的vout index 为 -1
	txin := TXInput{[]byte{}, -1, data}
	// 将奖励给to
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	// 找到足够的未花费输出
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}
		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}