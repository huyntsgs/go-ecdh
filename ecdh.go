package ecdh

import (
	"crypto"
	"io"
)

// The main interface for ECDH key exchange.
type ECDH interface {
	GenerateKey(io.Reader) (crypto.PrivateKey, crypto.PublicKey, error)
	Marshal(crypto.PublicKey) []byte
	MarshalPrivateKey(crypto.PrivateKey) []byte
	Unmarshal([]byte) (crypto.PublicKey, bool)
	UnmarshalPrivateKey(data []byte) crypto.PrivateKey
	GenerateSharedSecret(crypto.PrivateKey, crypto.PublicKey) ([]byte, error)
	GenSharedSecret32(privKey crypto.PrivateKey, pubKey crypto.PublicKey) ([]byte, error)
}
