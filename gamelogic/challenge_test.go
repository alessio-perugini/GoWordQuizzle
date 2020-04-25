package gamelogic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenges_AddChallenge(t *testing.T) {
	challenges := GetInstanceChallenges()
	c := Challenge{}
	if _, err := challenges.AddChallenge(c); err != nil {
		assert.EqualError(t, err, "illegal argument")
	}

	//Test insert of challenge
	c = Challenge{
		matchUser:    "Gino",
		matchFriend:  "Pino",
		id:           350,
		wordsToGuess: []map[string]string{{"cena": "dinner"}},
	}
	if v, err := challenges.AddChallenge(c); err == nil {
		assert.Equal(t, v, c)
	}

	if _, err := challenges.AddChallenge(c); err == nil {
		assert.EqualError(t, err, "challenge already added")
	}
}

func TestChallenges_RemoveChallenge(t *testing.T) {
	challenges := GetInstanceChallenges()
	c := Challenge{}
	if _, err := challenges.RemoveChallenge(c); err != nil {
		assert.EqualError(t, err, "illegal argument")
	}

	c = Challenge{
		matchUser:    "Gino",
		matchFriend:  "Pino",
		id:           350,
		wordsToGuess: []map[string]string{{"cena": "dinner"}},
	}
	//Test empty challenges slice
	if _, err := challenges.RemoveChallenge(c); err != nil {
		assert.EqualError(t, err, "challenge already removed")
	}

	challenges.AddChallenge(c)
	//Test successfully remove
	if v, err := challenges.RemoveChallenge(c); err == nil {
		assert.True(t, v)
	}
	//Test already removed challenge from the list
	if _, err := challenges.RemoveChallenge(c); err != nil {
		assert.EqualError(t, err, "challenge already removed")
	}
}

func TestChallenges_get(t *testing.T) {

}
