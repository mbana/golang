package swagger

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestParser_AccountInfoSwagger(t *testing.T) {
	require := require.New(t)

	swaggerPath := "../specifications/read-write/v3.1.1/account-info-swagger.yaml"
	logger := nullLogger()

	nonManadatoryFields, err := ParseSchema(swaggerPath, logger)
	require.NoError(err)
	require.NotNil(nonManadatoryFields)
}

func nullLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	return logger.WithFields(logrus.Fields{
		"test": "github.com/banaio/golang/openbanking_tools/swagger",
	}).Logger
}
