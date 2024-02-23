package crypto

func Ciphers(main Cipher, many ...Cipher) Cipher {
	if main == nil {
		panic("crypto.Ciphers: main cannot be nil")
	}
	for _, cipher := range many {
		if cipher == nil {
			panic("crypto.Ciphers: nil is not allowed")
		}
	}
	return ciphers{main, many}
}

type ciphers struct {
	main Cipher
	many []Cipher
}

func (c ciphers) Version() string {
	return c.main.Version()
}

func (c ciphers) Encrypt(plaintext []byte) ([]byte, error) {
	return c.main.Encrypt(plaintext)
}

func (c ciphers) Decrypt(version string, ciphertext[]byte,)([]byte ,error) {
	if c.main.Version() == version {
		return c.main.Decrypt(version, ciphertext)
	}
	for _, cipher := range c.many {
		if cipher.Version() == version {
			return cipher.Decrypt(version, ciphertext)
		}
	}
	return nil, InvalidVersion()
}

