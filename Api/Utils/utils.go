package utils

func Includes(array []string, search string) bool {
	for _, element := range array {
		if element == search {
			return true
		}
	}

	return false
}
