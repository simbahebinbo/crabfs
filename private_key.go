package crabfs

import (
	"bytes"
	"crypto/rand"
	crabfsCrypto "github.com/simbahebinbo/crabfs/crypto"
	"io"
)

// GenerateKeyPair generates a private and public keys ready to be used
func GenerateKeyPair() (crabfsCrypto.PrivKey, error) {
	privateKey, err := crabfsCrypto.PrivateKeyNew(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return privateKey.(crabfsCrypto.PrivKey), nil
}

// GenerateKeyPairReader generates a private and public keys ready to be used
func GenerateKeyPairReader() (io.Reader, error) {
	privateKey, err := GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	data, err := privateKey.Marshal()
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	return reader, nil
}

// ReadPrivateKey reads the key from a reader
func ReadPrivateKey(in io.Reader) (crabfsCrypto.PrivKey, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	key, err := crabfsCrypto.UnmarshalPrivateKey(data)
	if err != nil {
		return nil, err
	}

	return key, nil
}
