package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

const (
	checksumLength = 4
	version        = byte(0x00)
)

func ValidateAddress(address string) bool {
	pubhashKey := Base58Decode([]byte(address))
	actualCheckSum := pubhashKey[len(pubhashKey)-checksumLength:]
	version := pubhashKey[0]
	pubhashKey = pubhashKey[1 : len(pubhashKey)-checksumLength]

	targetcheckSum := CheckSum(append([]byte{version}, pubhashKey...))

	return bytes.Equal(actualCheckSum, targetcheckSum)
}

// func Base58Encode(input []byte) []byte {
// 	encode := base58.Encode(input)
// 	return []byte(encode)
// }

// func Base58Decode(input []byte) []byte {
// 	decode, err := base58.Decode(string(input[:]))

// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return decode
// }

func (w Wallet) Address() []byte {
	pubhash := PublicKeyHash(w.PublicKey)
	versionhash := append([]byte{version}, pubhash...)

	checksum1 := CheckSum(versionhash)

	fullHash := append(versionhash, checksum1...)
	address := Base58Encode(fullHash)

	// fmt.Printf("the public key : %x\n", w.PublicKey)
	// fmt.Printf("the public hash : %x\n", pubhash)
	// fmt.Printf("the Address : %x\n", address)

	return address

}
func PublicKeyHash(pubkey []byte) []byte {
	pubhash := sha256.Sum256(pubkey)

	hasher := ripemd160.New()

	_, err := hasher.Write(pubhash[:])
	if err != nil {
		log.Panic(err)

	}
	publicRipMD := hasher.Sum(nil)
	return publicRipMD
}

func CheckSum(payload []byte) []byte {
	firsthash := sha256.Sum256(payload)
	secondhash := sha256.Sum256(firsthash[:])

	return secondhash[:checksumLength]
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet

}
