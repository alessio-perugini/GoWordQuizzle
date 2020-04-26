package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendHttpRequest(t *testing.T) {
	body := SendHttpRequest("mela")
	_, err := UnmarshalMyMemoryAPI(body)
	assert.NoError(t, err)
}

func TestGetKeysFromMap(t *testing.T) {
	input := map[string]string{"a": "a1", "b": "b1", "c": "c1"}
	output := GetKeysFromMap(input)
	assert.Equal(t, output[0], "a", output[1], "b", output[2], "c")
}

func TestGetTranslationFromHttpGet(t *testing.T) {
	body := SendHttpRequest("banana")
	translatedWord := GetTranslationFromHttpGet(body)
	assert.Equal(t, translatedWord, "banana")
}
