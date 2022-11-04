package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

type Pair struct {
	left interface{}
	right interface{}
}

func main() {
	homework := readHomework("input.txt")
	fmt.Println(homework)
}

func readHomework(path string) (data []*Pair) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return
}

func parseLine(line string) *Pair {
	var pairs []*Pair
	for k, b := range line {
		switch s := string(b); s {
			case '[':
				p := new(Pair)
				pairs = append(pairs, p)
				if k == 0 {
					continue
				}
				e := string(line[k-1])
				if e == '[' {
					pairs[len(pairs)-2].(*Pair).left = p
				} else {
					// last rune is ','
					for j:=len(pairs)-1; j>=0; j-- {
						t, isPair := pairs[j].(*Pair)
						if isPair && t.right == nil {
							t.right = p
						}
					}
				}
			case ']':
			case ',':
				e := string(line[k-1])
				if e == ']' {
					continue
				}
				i := int(strconv.ParseInt(e, 10, 64))
				for j:=len(pairs)-1; j>=0; j-- {
					p, isPair := pairs[j].(*Pair)
					if isPair && p.left == nil {
						p.left = i
						break
					}
				}
			default:

		}
	}
}
