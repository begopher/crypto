package crypto

type Cipher interface {
	Version() string
	Encrypt([]byte) ([]byte, error)
	Decrypt(string, []byte) ([]byte ,error)
}
