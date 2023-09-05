package rsa_crypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"hash"
	"io"
	"os"
)

// https://pkg.go.dev/crypto/rsa
// https://stackoverflow.com/questions/62348923/rs256-message-too-long-for-rsa-public-key-size-error-signing-jwt
// https://stackoverflow.com/questions/11822607/what-is-the-rsa-max-block-size-to-encode
// https://stackoverflow.com/questions/14404757/how-to-encrypt-and-decrypt-plain-text-with-a-rsa-keys-in-go
// https://stackoverflow.com/questions/70718821/go-rsa-load-public-key

func EncryptOAEP(hash hash.Hash, random io.Reader, public *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := public.Size() - 2*hash.Size() - 2
	var encryptedBytesList [][]byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		encryptedBlockBytes, err := rsa.EncryptOAEP(hash, random, public, msg[start:finish], label)
		if err != nil {
			return nil, err
		}

		encryptedBytesList = append(encryptedBytesList, encryptedBlockBytes)
	}

	return bytes.Join(encryptedBytesList, nil), nil
}

func DecryptOAEP(hash hash.Hash, random io.Reader, private *rsa.PrivateKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := private.PublicKey.Size()
	var encryptedBytesList [][]byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, private, msg[start:finish], label)
		if err != nil {
			return nil, err
		}

		encryptedBytesList = append(encryptedBytesList, decryptedBlockBytes)
	}

	return bytes.Join(encryptedBytesList, nil), nil
}

func Usage() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from GenerateKey: %s\n", err)
		return
	}

	secretMessage := []byte("message to be encrypted")
	label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	ciphertext, err := EncryptOAEP(sha256.New(), rng, &privateKey.PublicKey, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	//=================================================================================================

	plaintext, err := DecryptOAEP(sha256.New(), nil, privateKey, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))

	// Remember that encryption only provides confidentiality. The
	// ciphertext should be signed before authenticity is assumed and, even
	// then, consider that messages might be reordered.

	pp.Println(string(plaintext) == string(secretMessage))
	//=================================================================================================

	message := []byte("message to be signed")

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPSS(rng, privateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature)
	//=================================================================================================

	err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, hashed[:], signature, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}

	// signature is a valid signature of message from the public key.

	pp.Println("signature verified")
	//=================================================================================================
}
