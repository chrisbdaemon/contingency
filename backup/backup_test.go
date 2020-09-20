package backup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileList(t *testing.T) {
	f, err := fileList(".")
	assert.Nil(t, err)
	assert.Greater(t, len(f), 0)

	f, err = fileList("doesntexist")
	assert.Error(t, err)
	assert.Equal(t, len(f), 0)
}
