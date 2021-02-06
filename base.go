package convert

import "strconv"

func intToString(i int) string {
	return strconv.Itoa(i)
}

func stringToInt(s string) int  {
	n, _ := strconv.Atoi(s)
	return n
}
