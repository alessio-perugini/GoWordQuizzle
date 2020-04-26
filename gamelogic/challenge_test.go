package gamelogic

import (
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChallenges_AddChallenge(t *testing.T) {
	challenges := GetInstanceChallenges()
	c := Challenge{}
	if _, err := challenges.AddChallenge(c); err != nil {
		assert.EqualError(t, err, "illegal argument")
	}

	u1 := common.User{}
	u2 := common.User{}
	u1.SetNickname("Gino")
	u2.SetNickname("Pino")

	m1 := Match{player: u1}
	m2 := Match{player: u2}

	//Test insert of challenge
	c = Challenge{
		matchUser:    m1,
		matchFriend:  m2,
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

	u1 := common.User{}
	u2 := common.User{}
	u1.SetNickname("Gino")
	u2.SetNickname("Pino")

	m1 := Match{player: u1}
	m2 := Match{player: u2}

	c = Challenge{
		matchUser:    m1,
		matchFriend:  m2,
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

func TestChallenge_generateWordsTranslation(t *testing.T) {
	input := []map[string]string{{"cocco": ""}, {"banana": ""}, {"limone": ""}}
	generateWordsTranslation(&input)
	expect := []map[string]string{{"cocco": "coconut"}, {"banana": "banana"}, {"limone": "lemon"}}
	assert.Equal(t, input, expect)
}
