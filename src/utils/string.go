package utils

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
