package backup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSize(t *testing.T) {
	testSizes := map[string]int64{
		"4.7GB": 5046586572,
		"1":     1,
		"5P":    5629499534213120,
	}

	for k, v := range testSizes {
		size, err := ParseSize(k)
		assert.Nil(t, err)
		assert.Equal(t, v, size)
	}

	_, err := ParseSize("-100")
	assert.Error(t, err)

	_, err = ParseSize("A")
	assert.Error(t, err)
}
