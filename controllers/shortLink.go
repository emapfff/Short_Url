package controllers

import (
	"crypto/rand"
	"math/big"
)

const (
	linkLength = 10
	charSet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func GenerateShortLink() string {
	b := make([]byte, linkLength)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			panic(err)
		}
		b[i] = charSet[num.Int64()]
	}
	return string(b)
}
