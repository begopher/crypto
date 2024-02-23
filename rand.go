package crypto

import(
	"io"
	"crypto/rand"
)

func Rand16() (bytes [16]byte, err error) {
	_, err = io.ReadFull(rand.Reader, bytes[:])
	return 
}

func Rand24() (bytes [24]byte, err error) {
	_, err = io.ReadFull(rand.Reader, bytes[:])
	return 
}

func Rand32() (bytes [32]byte, err error) {
	_, err = io.ReadFull(rand.Reader, bytes[:])
	return 
}
