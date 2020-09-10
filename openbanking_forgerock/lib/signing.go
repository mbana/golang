package lib

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// Key ...
type Key struct {
	PublicJWK          string
	PrivateJWK         string
	PublicCertificate  *rsa.PublicKey
	PrivateCertificate *rsa.PrivateKey
}

// Keys ...
type Keys struct {
	Encryption Key
	Signature  Key
	Transport  Key
}

// Sign ...
func Sign(claims jwt.Claims, kid string) string {
	// TODO: move
	// software_statement and aud
	// iat, exp and jti into this function.
	//
	// now := time.Now()
	// iat := now.Unix()
	// // exp := now.Add(24 * time.Hour).Unix()
	// exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
	// jti := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	if kid != "" {
		token.Header["kid"] = kid
	}

	keys := GetKeys()
	signedString, err := token.SignedString(keys.Signature.PrivateCertificate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"signedString": signedString,
			"err":          err,
		}).Fatal("Sign")
	}

	return signedString
}

// GetKeys ...
func GetKeys() *Keys {
	signature := getKey("signature")
	return &Keys{
		Signature: signature,
	}
}

func readKey(keyType string, publicKey bool) []byte {
	extension := "key"
	if publicKey {
		extension = "pem"
	}

	path := fmt.Sprintf(".ignore/keys/%s/d6c3f49c-7112-4c5c-9c9d-84926e992c74.%s", keyType, extension)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("readKey")
	}

	logrus.WithFields(logrus.Fields{
		"path": path,
	}).Debug("readKey")

	return file
}

func getKey(keyType string) Key {
	privateCertificate, err := jwt.ParseRSAPrivateKeyFromPEM(readKey(keyType, false))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("getKeys")
	}

	publicCertificate, err := jwt.ParseRSAPublicKeyFromPEM(readKey(keyType, true))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Fatal("getKeys")
	}

	return Key{
		PublicCertificate:  publicCertificate,
		PrivateCertificate: privateCertificate,
	}
}
