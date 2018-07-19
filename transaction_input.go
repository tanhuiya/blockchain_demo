package main

import "bytes"

type TXInput struct {
	// 一个交易输入引用了之前一笔交易的输出，txid 表示引用了之前哪笔交易
	Txid []byte
	// 引用之前交易输出的 index
	Vout int
	// 用来解锁 vout 的数据
	Signature []byte

	PubKey []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
