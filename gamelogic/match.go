package gamelogic

import (
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"github.com/tevino/abool"
	"time"
)

type Match struct {
	player       common.User
	startTime    time.Time
	endTime      time.Time
	isEnd        *abool.AtomicBool
	wordsToGuess []map[string]string
	totalWords   int
	wrongWords   int
	rightWords   int
	notAnswerd   int
	finalScore   int
	//TODO Qualcosa relativa al socket?
}

func NewMatch(u common.User, c Challenge) *Match {
	startTime := time.Now()
	return &Match{
		wordsToGuess: c.wordsToGuess,
		player:       u,
		isEnd:        abool.New(),
		totalWords:   len(c.wordsToGuess),
		startTime:    startTime,
		endTime:      startTime.Add(common.ParseDuration(common.MATCH_MAX_TIME)),
	}
}

func (m *Match) Play() {

}
