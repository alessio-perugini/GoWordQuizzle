package common

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetTranslationFromHttpGet(body []byte) string {
	json, _ := UnmarshalMyMemoryAPI(body)
	return json.ResponseData.TranslatedText
}

func SendHttpRequest(word string) []byte {
	resp, err := http.Get(API_URL + "get?q=" + word + "&langpair=it|en")
	if err != nil {
		log.Println(err)
		return make([]byte, 0)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return make([]byte, 0)
	}

	return body
}

func GetKeysFromMap(v map[string]string) []string {
	keys := make([]string, len(v))

	i := 0
	for k := range v {
		keys[i] = k
		i++
	}

	return keys
}
