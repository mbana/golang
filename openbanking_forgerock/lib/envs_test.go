package lib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnvs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	defer func() {
		os.Unsetenv("SSID")
		os.Unsetenv("ORGANISATION_ID")
		os.Unsetenv("SSA")
		os.Unsetenv("KID")
	}()

	SSID := "SSID"
	organisationID := "ORGANISATION_ID"
	SSA := "SSA"
	KID := "KID"

	os.Setenv("SSID", SSID)
	os.Setenv("ORGANISATION_ID", organisationID)
	os.Setenv("SSA", SSA)
	os.Setenv("KID", KID)

	envs, err := GetEnvs()
	assert.NoError(err)
	assert.Equal(SSID, envs.SSID)
	assert.Equal(organisationID, envs.OrganisationID)
	assert.Equal(SSA, envs.SSA)
	assert.Equal(KID, envs.KID)
}

func Test_GetEnvsError(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	defer func() {
		os.Unsetenv("SSID")
		os.Unsetenv("ORGANISATION_ID")
		os.Unsetenv("SSA")
		os.Unsetenv("KID")
	}()

	if envs, err := GetEnvs(); assert.Error(err) {
		assert.Nil(envs)
		assert.Error(err)
		assert.Equal("missing Software Statement ID environmental variable: SSID=<value>", err.Error())
	}

	os.Setenv("SSID", "SSID")
	if envs, err := GetEnvs(); assert.Error(err) {
		assert.Nil(envs)
		assert.Error(err)
		assert.Equal("missing Organisation ID environmental variable: ORGANISATION_ID=<value>", err.Error())
	}

	os.Setenv("ORGANISATION_ID", "ORGANISATION_ID")
	if envs, err := GetEnvs(); assert.Error(err) {
		assert.Nil(envs)
		assert.Error(err)
		assert.Equal("missing Software Statement Assertion environmental variable: SSA=<value>", err.Error())
	}

	os.Setenv("SSA", "SSA")
	if envs, err := GetEnvs(); assert.Error(err) {
		assert.Nil(envs)
		assert.Error(err)
		assert.Equal("missing Key ID environmental variable: KID=<value>", err.Error())
	}
}
