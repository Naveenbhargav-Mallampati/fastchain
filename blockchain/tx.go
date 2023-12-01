package blockchain

import (
	"bytes"
	"encoding/gob"

	"fastchain.com/corechain/wallet"
)

type TxInp struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

type TxOutputs struct {
	Outputs []TxOut
}

type TxOut struct {
	Value      int
	PubKeyHash []byte
}

// func (tx *TxInp) CanUnLock(data string) bool {
// 	return tx.Sig == data
// }

// func (tx *TxOut) CanBeUnLocked(data string) bool {
// 	return tx.PubKey == data
// }

func (in *TxInp) UsesKey(pubkeyhash []byte) bool {
	lockinghash := wallet.PublicKeyHash(pubkeyhash)

	return bytes.Equal(lockinghash, pubkeyhash)
}

func (out *TxOut) Lock(address []byte) {
	pubHashKey := wallet.Base58Decode(address)
	pubHashKey = pubHashKey[1 : len(pubHashKey)-4]
	out.PubKeyHash = pubHashKey
}

func (out *TxOut) IsLockedWithKey(pubhashkey []byte) bool {
	return bytes.Equal([]byte(out.PubKeyHash), pubhashkey)
}

func NewTxOut(value int, address string) *TxOut {
	txo := &TxOut{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (outs TxOutputs) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(outs)
	Handle(err)
	return buffer.Bytes()

}

func DeSerializeOutputs(data []byte) TxOutputs {
	var outputs TxOutputs
	decode := gob.NewDecoder(bytes.NewReader(data))
	err := decode.Decode(&outputs)
	Handle(err)
	return outputs

}
