package signer

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"

	"github.com/banaio/openbankingforgerock/config"

	"strings"
)

const (
	applicationID       = "5cc9660469b9ba001a381708"
	organisationID      = "5b507065b093465496d238a8"
	softwareStatementID = "5cc9660355e1530017c48a8d"
	issuerID            = "4c625555-2e29-414a-b6e0-f8cb78256fdf"
)

func MakeJWSSignature(config *config.Config, keys *Keys, payload []byte) (string, error) {
	privateKey := keys.Signature.PrivateCertificate

	now := time.Now()
	iat := now.Unix()
	crit := []string{
		"b64",
		"http://openbanking.org.uk/iat",
		"http://openbanking.org.uk/tan",
		"http://openbanking.org.uk/iss",
	}
	iss := config.ClientID

	opts := (&jose.SignerOptions{
		ExtraHeaders: map[jose.HeaderKey]interface{}{
			"kid":                           config.KID,
			"b64":                           false,
			"crit":                          crit,
			"http://openbanking.org.uk/tan": "openbanking.org.uk",
			"http://openbanking.org.uk/iat": iat,
			"http://openbanking.org.uk/iss": iss,
			jose.HeaderContentType:          "application/json",
			jose.HeaderType:                 "JOSE",
		},
	})

	// Instantiate a signer using RSASSA-PSS (SHA256) with the given private key.
	signer, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.PS256,
		Key:       privateKey,
	}, opts)
	if err != nil {
		return "", errors.Wrap(err, "makeJWSSignature on jose.NewSigner")
	}

	// Sign a sample payload. Calling the signer returns a protected JWS object,
	// which can then be serialized for output afterwards. An error would
	// indicate a problem in an underlying cryptographic primitive.
	// var payload = []byte("Lorem ipsum dolor sit amet")
	object, err := signer.Sign(payload)
	if err != nil {
		return "", errors.Wrap(err, "makeJWSSignature on signer.Sign")
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized, err := object.CompactSerialize()
	if err != nil {
		return "", errors.Wrap(err, "makeJWSSignature on object.CompactSerialize")
	}

	// Parse the serialized, protected JWS object. An error would indicate that
	// the given input did not represent a valid message.
	object, err = jose.ParseSigned(serialized)
	if err != nil {
		return "", errors.Wrap(err, "failed jose.ParseSigned")
	}

	// // Now we can verify the signature on the payload. An error here would
	// // indicate the the message failed to verify, e.g. because the signature was
	// // broken or the message was tampered with.
	// if _, err := object.Verify(&privateKey.PublicKey); err != nil {
	// 	return "", errors.Wrap(err, "failed object.Verify")
	// }

	// Now we can verify the signature on the payload. An error here would
	// indicate the the message failed to verify, e.g. because the signature was
	// broken or the message was tampered with.
	if err := object.DetachedVerify(payload, &privateKey.PublicKey); err != nil {
		return "", errors.Wrap(err, "failed object.DetachedVerify")
	}

	splittedJWS := strings.Split(serialized, ".")
	if len(splittedJWS) != 3 {
		return "", errors.Errorf("makeJWSSignature on len(splittedJWS) != 3, splittedJWS=%#v, serialized=%+v", splittedJWS, serialized)
	}
	signature := strings.Join([]string{
		splittedJWS[0],
		"",
		splittedJWS[2],
	}, ".")
	// signature := splittedJWS[0] + "." + "." + splittedJWS[2]

	fmt.Printf("%+v\n", strings.Repeat("-", 50))
	fmt.Printf("  makeJWSSignature signature:\n%+v\n", signature)
	fmt.Printf("%+v\n", strings.Repeat("-", 50))

	return signature, nil
}
