package gamelogic

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetInstanceDictionary(t *testing.T) {
	d := GetInstanceDictionary()
	n := 3

	assert.NotEmpty(t, d.words)
	if _, err := d.GetNwords(math.MaxInt64); err != nil {
		assert.EqualError(t, err, "dictionary contains less words")
	}

	if v, err := d.GetNwords(n); err == nil {
		assert.Len(t, v, n)
	}
}
