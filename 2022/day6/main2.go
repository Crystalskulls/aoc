package main

import (
	"fmt"
	"os"
)

type Handheld struct {
	ComSystem *ComSystem
}

type ComSystem struct {
	Signal               []byte
	StartOfPacketMarker  int
	StartOfMessageMarker int
}

func main() {
	handheld := new(Handheld)
	handheld.ComSystem = new(ComSystem)
	handheld.ComSystem.receiveSignal()
	handheld.ComSystem.detectStartOfPacketMarker()
	handheld.ComSystem.detectStartOfMessageMarker()
	fmt.Printf("Part One: %d - Part Two: %d\n", handheld.ComSystem.StartOfPacketMarker, handheld.ComSystem.StartOfMessageMarker)
}

func (cs *ComSystem) receiveSignal() {
	cs.Signal, _ = os.ReadFile("input.txt")
}

func (cs *ComSystem) detectStartOfPacketMarker() {
	cs.StartOfPacketMarker = detectMarker(cs.Signal, 4)
}

func (cs *ComSystem) detectStartOfMessageMarker() {
	cs.StartOfMessageMarker = detectMarker(cs.Signal, 14)
}

func detectMarker(signal []byte, distChars int) (marker int) {
	cm := make(map[byte]struct{})
	for i := distChars - 1; i <= len(signal); i++ {
		cs := signal[i-(distChars-1) : i+1]
		for _, c := range cs {
			cm[c] = struct{}{}
		}
		if len(cm) == distChars {
			marker = i + 1
			break
		}
		cm = make(map[byte]struct{})
	}
	return
}
