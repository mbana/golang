package signer

import (
	"crypto/rsa"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"github.com/banaio/openbankingforgerock/config"
)

// Signer -
type Signer struct {
	Keys   *Keys
	Config *config.Config
}

// NewSigner -
func NewSigner(keys *Keys, config *config.Config) (*Signer, error) {
	if keys == nil {
		return nil, errors.New("keys == nil: NewSigner")
	}
	if config == nil {
		return nil, errors.New("config == nil: NewSigner")
	}
	return &Signer{
		Keys:   keys,
		Config: config,
	}, nil
}

// Sign ...
func (s *Signer) Sign(claims jwt.Claims) string {
	var fixedSigningMethodPS256 = &jwt.SigningMethodRSAPSS{
		SigningMethodRSA: jwt.SigningMethodPS256.SigningMethodRSA,
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
	}
	// TODO: move
	// software_statement and aud
	// iat, exp and jti into this function.
	//
	// now := time.Now()
	// iat := now.Unix()
	// // exp := now.Add(24 * time.Hour).Unix()
	// exp := time.Date(2019, 03, 29, 0, 0, 0, 0, time.UTC).Unix()
	// jti := uuid.New().String()
	// token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)

	// TODO(mbana): Switch on algorithm type.
	alg := fixedSigningMethodPS256

	token := jwt.NewWithClaims(alg, claims)
	kid := s.Config.KID
	if kid != "" {
		token.Header["kid"] = kid
	}

	keys := s.Keys
	signedString, err := token.SignedString(keys.Signature.PrivateCertificate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"signedString": signedString,
			"err":          err,
		}).Fatal("Sign")
	}

	return signedString
}
