package test

import(
	"testing"
	//"slices"
	"github.com/begopher/crypto"
)


func Test_AES128GCM(t *testing.T) {
	key, err := crypto.Rand16()
	if err != nil {
		t.Fatal(err)
	}
	version := "first version"
	cipher, err := crypto.AES128GCM(version, key)
	if err != nil {
		t.Fatal(err)
	}
	plaintext := "gopher"
	ciphertext, err := cipher.Encrypt([]byte(plaintext))
	if err != nil {
		t.Fatal(err)
	}
	if string(ciphertext) ==  plaintext {
		t.Errorf("Ciphertext should not be equal to plaintext")
	}
	plain, err := cipher.Decrypt(cipher.Version(), ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	expected := plaintext
	if got := string(plain); got != expected {
		t.Errorf("Expected plaintext is (%v) got (%v)", expected, got)
	}
	
}
