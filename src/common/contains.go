package common

func Contains(array []int, item int) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func Containsstr(array string, item string) bool {
	for _, v := range array {
		if string(v) == item {
			return true
		}
	}
	return false
}

func ContainsRune(word string, letter string) bool {
    for _, runeValue := range word {
        if string(runeValue) == letter {
            return true
        }
    }
    return false
}
