package lib

import (
	"testing"

	"github.com/banaio/golang/project_template/lib/utils"
)

func Test_Lib_Template(t *testing.T) {
	assert, require, logger := utils.MakeTestCaseDependencies(t)
	_, _, _ = assert, require, logger
	t.Skipf("test=%+v is intentionally ignored it serves as a template to copy when making another test", t.Name())

	require.Equal("true", "true")
	assert.Equal("true", "true")
}
