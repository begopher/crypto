package crypto

import(
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func AES128GCM(version string, key [16]byte) (Cipher, error) {
	return aesGCM(version, key[:])
}

func AES192GCM(version string, key [24]byte) (Cipher, error) {
	return aesGCM(version, key[:])
}

func AES256GCM(version string, key [32]byte) (Cipher, error) {
	return aesGCM(version, key[:])
}

func aesGCM(version string, key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return _aes{version, aead}, nil
}

type _aes struct {
	version string
	aead    cipher.AEAD 
}

func (a _aes) Version() string {
	return a.version
}

func (a _aes) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, a.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return a.aead.Seal(nonce, nonce, plaintext, nil), nil
}

func (a _aes) Decrypt(version string, ciphertext[]byte)([]byte ,error) {
	if a.Version() != version {
		return nil, InvalidVersion()
	}
	return a.aead.Open(
		nil,
		ciphertext[:a.aead.NonceSize()],
		ciphertext[a.aead.NonceSize():],
		nil,
	)
}
