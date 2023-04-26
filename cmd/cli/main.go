package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var notes = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
}

var accidentals = []string{
	"",
	"#",
	"b",
}

const usage = "randomnotes [number of notes]"

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("main: invalid arguments. usage: %s", usage)
	}

	noteCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("main: error parsing note count: %v", err)
	}

	currentCount := 0
	selectedNotes := make(map[string]bool)

	for currentCount < noteCount {
		selectedNote := notes[rand.Intn(len(notes))] + accidentals[rand.Intn(len(accidentals))]

		if _, ok := selectedNotes[selectedNote]; ok {
			continue
		}

		selectedNotes[selectedNote] = true
		currentCount++
	}

	for k, _ := range selectedNotes {
		fmt.Println(k)
	}

}
