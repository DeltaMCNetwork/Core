package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

const (
	BitSize = 1024
)

type Keypair struct {
	Public  *PublicKey
	Private *rsa.PrivateKey
}

type PublicKey struct {
	Len int
	Key []byte
}

func NewKeypair() *Keypair {
	privateKey, err := rsa.GenerateKey(rand.Reader, BitSize)
	if err != nil {
		panic("error generting private key " + err.Error())
	}

	publicKey := &privateKey.PublicKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic("error marshaling public key " + err.Error())
	}

	return &Keypair{
		Public:  &PublicKey{Len: len(publicKeyBytes), Key: publicKeyBytes},
		Private: privateKey,
	}
}

func (k *Keypair) Decrypt(bytearr []byte) ([]byte, error) {
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, k.Private, bytearr) //

	fmt.Printf("Length of decrypted is %d\n", len(decrypted))

	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
