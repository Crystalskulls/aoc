package flag

import "flag"

func CreatePuzzleFlag() (puzzleFlag int) {
	flag.IntVar(&puzzleFlag, "puzzle", 1, "Which puzzle should be solved? Select '1' or '2'")
	flag.IntVar(&puzzleFlag, "p", 1, "Which puzzle should be solved? Select '1' or '2'")
	flag.Parse()
	return
}
