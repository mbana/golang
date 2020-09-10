package cli

import (
	"testing"

	"github.com/banaio/golang/project_template/lib/utils"
)

func Test_CLI_Template(t *testing.T) {
	assert, require := utils.MakeAssertAndRequire(t)
	_, _ = assert, require
	t.Skipf("test=%+v is intentionally ignored it serves as a template to copy when making another test", t.Name())

	require.Equal(true, false)
	assert.Equal(true, false)
}
