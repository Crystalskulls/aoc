package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a/"
	i := strings.LastIndex(path[:len(path)-1], "/")
	parent := path[:i+1]
	fmt.Println(i)
	fmt.Println(parent)
}
