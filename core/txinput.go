package core

import "bytes"

type TXInput struct {
	Txid      []byte // Transaction ID
	Vout      int    // TODO: float32
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
