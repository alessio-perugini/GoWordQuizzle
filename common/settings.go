package common

const (
	RPC_PORT            uint16 = 50000
	TCP_PORT            uint16 = 50001
	UDP_PORT            uint16 = 50002
	UDP_TIMEOUT         string = "2s"
	MATCH_MAX_TIME             = "60s"
	DICTIONARY_FILE            = "../data/dictionary.txt"
	HOST_NAME                  = "localhost"
	API_URL                    = "https://api.mymemory.translated.net/"
	JSON_FILE                  = "user.json"
	MAX_WORD_TO_GEN            = 10
	MIN_WORD_TO_GEN            = 3
	EXTRA_POINTS               = 3
	AUTO_SAVE_JSON             = false
	AUTO_SAVE_FREQUENCY        = "20s"
)
