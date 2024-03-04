package utils

import "encoding/json"

func Includes(array []string, search string) bool {
	for _, element := range array {
		if element == search {
			return true
		}
	}

	return false
}

// takes and input entity and output json format
func ApplyContext(input interface{}, context interface{}) {
	// convert input to json
	jsonInput, _ := json.Marshal(input)

	// convert json to struct
	json.Unmarshal(jsonInput, context)
}
