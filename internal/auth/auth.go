package auth

import (
	"arikatto/internal/config"
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadCertificates function
// Public function that searches and loads the certificates
// of the keys for the signature of the jwt.
// Params:
// - privateFile string
// - publicFile string
// Return:
// - err error
func LoadCertificates() {
	var err error = nil

	configurationReader := config.Get()

	once.Do(func() {
		err = loadFiles(configurationReader.GetPrivatePath(), configurationReader.GetPublicPath())
	})

	if err != nil {
		fmt.Println("--> Certificates could not be loaded")
		log.Fatal(err)
	}
}

// loadFiles function
// Private function that obtains the location paths of
// the certificates, reads and stores them in type
// []byte and then calls the parseRSA function.
// Params:
// - privateFile string
// - publicFile string
// Return:
// - err error
func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := os.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicBytes, err := os.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return parseRSA(privateBytes, publicBytes)
}

// parseRSA function
// Private function that receives the certificates in
// type []byte converts them to type rsa and stores them
// in the container variables signKey and verifyKey.
// Params:
// - privateBytes []byte
// - publicBytes []byte
// Return:
// - err error
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}

	return nil
}
