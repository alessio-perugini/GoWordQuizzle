package gamelogic

import (
	"bufio"
	"errors"
	"github.com/alessio-perugini/GoWordQuizzle/common"
	"log"
	"os"
	"sync"
)

type Dictionary struct {
	words []string //TODO check if is a concurrent part
}

var instanceDictionary *Dictionary
var once sync.Once

func GetInstanceDictionary() *Dictionary {
	once.Do(func() {
		instanceDictionary = &Dictionary{words: loadDictionaryFromFile(common.DICTIONARY_FILE)}
	})

	return instanceDictionary
}

func loadDictionaryFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	dict := make([]string, 0, 20)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return dict
}

func (d *Dictionary) GetNwords(n int) ([]map[string]string, error) {
	if n > len(d.words) {
		return make([]map[string]string, 0), errors.New("dictionary contains less words")
	}

	//TODO Convertire in struct?
	chosenWords := make([]map[string]string, 0, n)

	for i := 0; i < n; i++ { //TODO creare randomizzazione
		translation := map[string]string{d.words[i]: ""}
		chosenWords = append(chosenWords, translation)
	}

	return chosenWords, nil
}
