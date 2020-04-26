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
	isEnd        abool.AtomicBool //TODO lock
	wordsToGuess []map[string]string
	totalWords   int
	wrongWords   int
	rightWords   int
	notAnswerd   int
	finalScore   int
	//TODO Qualcosa relativa al socket?
}
