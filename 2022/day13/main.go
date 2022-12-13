package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	rightOrder *bool
}

func main() {
	fmt.Printf("PartOne: %d\n", findRightOrderedPairs())
	fmt.Printf("PartTwo: %d\n", findDecoderKey())
}

func findRightOrderedPairs() (sum int) {
	file, _ := os.ReadFile("input.txt")
	pairs := strings.Split(string(file), "\n\n")

	for i, pair := range pairs {
		var a, b []interface{}
		lines := strings.Split(pair, "\n")
		json.Unmarshal([]byte(lines[0]), &a)
		json.Unmarshal([]byte(lines[1]), &b)
		if *compare(a, b) {
			sum += i + 1
		}
	}
	return sum
}

func findDecoderKey() (decoderKey int) {
	decoderKey = 1
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	sortedPackets := make([][]interface{}, 0)
	var firstPacket []interface{}
	json.Unmarshal([]byte(lines[0]), &firstPacket)
	sortedPackets = append(sortedPackets, firstPacket)
	lines = lines[1:]
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var nextPacket []interface{}
		json.Unmarshal([]byte(line), &nextPacket)
		inserted := false
		for i, packet := range sortedPackets {
			if *compare(nextPacket, packet) {
				sortedPackets = insert(sortedPackets, i, nextPacket)
				inserted = true
				break
			}
		}
		if !inserted {
			sortedPackets = append(sortedPackets, nextPacket)
		}
	}
	for i, sortedPacket := range sortedPackets {
		s := fmt.Sprintf("%v", sortedPacket)
		if s == "[[2]]" {
			decoderKey *= i + 1
		} else if s == "[[6]]" {
			decoderKey *= i + 1
		}
	}
	return decoderKey
}

func insert(sortedPackets [][]interface{}, i int, packet []interface{}) [][]interface{} {
	if len(sortedPackets) == i {
		return append(sortedPackets, packet)
	}
	sortedPackets = append(sortedPackets[:i+1], sortedPackets[i:]...)
	sortedPackets[i] = packet
	return sortedPackets
}

func compare(a []interface{}, b []interface{}) *bool {
	for i, element := range a {
		if i > len(b)-1 {
			return new(bool)
		}
		av, aIsFloat := element.(float64)
		bv, bIsFloat := b[i].(float64)

		if aIsFloat && bIsFloat {
			if av > bv {
				return new(bool)
			} else if av < bv {
				b := true
				return &b
			}
			continue
		}

		asoi, aIsSliceOfInterfaces := element.([]interface{})
		bsoi, bIsSliceOfInterfaces := b[i].([]interface{})

		if aIsSliceOfInterfaces && bIsSliceOfInterfaces {
			rightOrder := compare(asoi, bsoi)
			if rightOrder != nil {
				return rightOrder
			}
		}

		if aIsSliceOfInterfaces && !bIsSliceOfInterfaces {
			k := make([]interface{}, 0)
			k = append(k, b[i].(float64))
			rightOrder := compare(asoi, k)
			if rightOrder != nil {
				return rightOrder
			}
		} else if !aIsSliceOfInterfaces && bIsSliceOfInterfaces {
			k := make([]interface{}, 0)
			k = append(k, element.(float64))
			rightOrder := compare(k, bsoi)
			if rightOrder != nil {
				return rightOrder
			}
		}

	}
	if len(a) < len(b) {
		b := true
		return &b
	}
	return nil
}
