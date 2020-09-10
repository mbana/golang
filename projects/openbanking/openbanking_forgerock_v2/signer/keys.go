package signer

import (
	"crypto/rsa"
	"io/ioutil"

	"crypto/tls"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/banaio/openbankingforgerock/config"
)

// Key ...
type Key struct {
	Algorithm          string
	PublicJWK          string
	PrivateJWK         string
	PublicCertificate  *rsa.PublicKey
	PrivateCertificate *rsa.PrivateKey
	Certificate        tls.Certificate
}

// Keys ...
type Keys struct {
	Encryption *Key // RSA-OAEP-256
	Signature  *Key // PS256
	Transport  *Key // PS256
}

// NewKeys -
func NewKeys(config *config.Config) (*Keys, error) {
	encryption, err := getKey(config.EncryptionPrivate, config.EncryptionPublic)
	if err != nil {
		return nil, err
	}
	signature, err := getKey(config.SignaturePrivate, config.SignaturePublic)
	if err != nil {
		return nil, err
	}
	transport, err := getKey(config.TransportPrivate, config.TransportPublic)
	if err != nil {
		return nil, err
	}

	return &Keys{
		Encryption: encryption,
		Signature:  signature,
		Transport:  transport,
	}, nil
}

func getKey(filenamePrivate string, filenamePublic string) (*Key, error) {
	privateKey, err := ioutil.ReadFile(filenamePrivate)
	if err != nil {
		return nil, err
	}
	privateCertificate, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := ioutil.ReadFile(filenamePublic)
	if err != nil {
		return nil, err
	}
	publicCertificate, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	certificate, err := tls.X509KeyPair(publicKey, privateKey)
	if err != nil {
		return nil, err
	}

	return &Key{
		PublicCertificate:  publicCertificate,
		PrivateCertificate: privateCertificate,
		Certificate:        certificate,
	}, nil
}
