package common

import (
	"encoding/json"
	"os"
)

func WriteObjToJSONFile(fname string, obj interface{}){
	file, _ := os.OpenFile(fname, os.O_CREATE, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(obj)
}

func ObjectToJSON(obj interface{}) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}

	return string(data)
}

func GetObjectFromJSONFile(fname string) interface{}{
	file, _ := os.Open(fname)
	defer file.Close()

	var data interface{}
	decoder := json.NewDecoder(file)
	decoder.Decode(&data)

	return data
}

func readJSONToken(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)

	var filteredData []map[string]interface{}

	// Read the array open bracket
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)

		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}