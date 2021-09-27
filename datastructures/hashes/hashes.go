package hashes

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

func DoHash() {
	var firstHash hash.Hash
	firstHash = sha256.New()
	fmt.Println(firstHash.Write([]byte("test encoding")))
}
