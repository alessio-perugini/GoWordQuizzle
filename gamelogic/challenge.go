package gamelogic

import (
	"errors"
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Challenge struct {
	matchUser    Match
	matchFriend  Match
	id           int
	wordsToGuess []map[string]string
}

type Challenges struct {
	listChallenges []map[int]Challenge //TODO add mutex
}

var instanceChallenges *Challenges
var onceChalenges sync.Once

func GetInstanceChallenges() *Challenges {
	onceChalenges.Do(func() {
		instanceChallenges = &Challenges{listChallenges: make([]map[int]Challenge, 0, 100)}
	})

	return instanceChallenges
}

//TODO neeed to sync
func (c *Challenges) AddChallenge(chllng Challenge) (Challenge, error) {
	if chllng.id == 0 {
		return Challenge{}, errors.New("illegal argument")
	}
	if i, _ := c.get(chllng); i != -1 {
		return Challenge{}, errors.New("challenge already added")
	}

	c.listChallenges = append(c.listChallenges, map[int]Challenge{chllng.id: chllng})

	return (c.listChallenges[len(c.listChallenges)-1])[chllng.id], nil
}

func (c *Challenges) RemoveChallenge(chllng Challenge) (bool, error) {
	if chllng.id == 0 {
		return false, errors.New("illegal argument")
	}
	i, _ := c.get(chllng)
	if len(c.listChallenges) == 0 || i == -1 {
		return false, errors.New("challenge already removed")
	}

	c.listChallenges = append(c.listChallenges[:i], c.listChallenges[i+1:]...)

	return true, nil
}

func (c *Challenges) get(v Challenge) (int, Challenge) {
	for i := 0; i < len(c.listChallenges); i++ {
		localChallenge := (c.listChallenges[i])[v.id]
		if localChallenge.id == v.id {
			return i, localChallenge
		}
	}

	return -1, Challenge{}
}

func NewChallenge(id int) *Challenge {
	rand.Seed(time.Now().UnixNano())
	maxWord := rand.Intn(common.MAX_WORD_TO_GEN + 1)
	kWordToSend := int(math.Max(float64(maxWord), common.MIN_WORD_TO_GEN))
	words, err := GetInstanceDictionary().GetNwords(kWordToSend)
	if err != nil {
		return &Challenge{}
	}

	generateWordsTranslation(&words)

	return &Challenge{
		id:           id,
		wordsToGuess: words,
	}
}

func (c *Challenge) SetMatch(user, friend Match) {
	if c.matchUser.player.GetNickname() != user.player.GetNickname() {
		c.matchUser = user
	}
	if c.matchUser.player.GetNickname() != friend.player.GetNickname() {
		c.matchFriend = friend
	}
}

func generateWordsTranslation(words *[]map[string]string) {
	for _, e := range *words {
		keys := common.GetKeysFromMap(e)
		e[keys[0]] = common.GetTranslationFromHttpGet(common.SendHttpRequest(keys[0]))
	}
}
