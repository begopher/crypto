package test

import(
	"testing"
	"github.com/begopher/crypto"
)

func Test_Ciphers(t *testing.T){
	oldVersion := "one"
	oldKey, err := crypto.Rand16()
	if err != nil {
		t.Fatal(err)
	}
	oldCipher, err := crypto.AES128GCM(oldVersion, oldKey)
	if err != nil {
		t.Fatal(err)
	}
	newVersion := "two"
	newKey, err := crypto.Rand32()
	if err != nil {
		t.Fatal(err)
	}
	newCipher, err := crypto.AES256GCM(newVersion, newKey)
	if err != nil {
		t.Fatal(err)
	}
	cipher := crypto.Ciphers(newCipher, oldCipher)
	plantext := "golang"
	ciphertext, err := oldCipher.Encrypt([]byte(plantext))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := cipher.Decrypt(oldCipher.Version(), ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(bytes); got != plantext {
		t.Errorf("Expected decrypted value is (%s) got (%s)", plantext, got)
	}
	ciphertext, err = cipher.Encrypt([]byte(plantext))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err = newCipher.Decrypt(cipher.Version(), ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(bytes); got != plantext {
		t.Errorf("Expected decrypted value is (%s) got (%s)", plantext, got)
	}
}
