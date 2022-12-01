package strconv

import (
	"log"
	"strconv"
)

func ParseToInt(n string) int {
	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse number %s to an Integer; err: %v\n", n, err)
	}
	return int(i)
}
