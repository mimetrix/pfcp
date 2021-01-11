package version_test

import (
	"testing"

	"github.com/mimetrix/pfcp/version"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-03-31-01", version.GetVersion())
}
