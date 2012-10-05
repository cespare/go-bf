package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	/*"strings"*/
	/*"time"*/
)

// Setup
var instructions []byte
var instrPos = 0
var tape = make([]byte, 30000)
var dataPos = 0

func main() {
	// Read the program code
	contents := make([]byte, 0)
	switch len(os.Args) {
	case 1:
		// Read from stdin
		input := bufio.NewReader(os.Stdin)
		for {
			part, err := input.ReadBytes('\n')
			if (err != nil) && (err != io.EOF) {
				log.Fatal("Error reading stdin.")
			}
			contents = append(contents, part...)
			if err == io.EOF {
				break
			}
		}
	case 2:
		// Read from file
		filename := os.Args[1]
		var err error
		contents, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Bad arguments.")
	}

	instructions = contents

	// Execution
	for {
		// Termination
		if instrPos >= len(instructions) {
			break
		}

		// Print state.
		/*leaderWidth := 15*/
		/*fmt.Printf("tape[0:200]:  (")*/
		/*for i := 0; i < 100 && i < len(tape); i++ {*/
			/*datum := tape[i]*/
			/*if datum == 0 {*/
				/*fmt.Printf("_")*/
			/*} else if datum < ' ' || datum > '~' {*/
				/*fmt.Printf("?")*/
			/*} else {*/
				/*fmt.Printf("%c", datum)*/
			/*}*/
		/*}*/
		/*fmt.Printf(")\n")*/
		/*fmt.Println(strings.Repeat(" ", dataPos+leaderWidth) + "^")*/

		/*fmt.Printf("instructions: (")*/
		/*for i := 0; i < 200 && i < len(instructions); i++ {*/
			/*instr := instructions[i]*/
			/*if instr < ' ' || instr > '~' {*/
				/*fmt.Printf(" ")*/
			/*} else {*/
				/*fmt.Printf("%c", instr)*/
			/*}*/
		/*}*/
		/*fmt.Printf(")\n")*/
		/*fmt.Println(strings.Repeat(" ", instrPos+leaderWidth) + "^")*/
		/*time.Sleep(100 * time.Millisecond)*/

		/*fmt.Printf("%c\n", contents[instrPos])*/
		datum := tape[dataPos]
		switch instructions[instrPos] {
		case '>':
			moveDataPos(1)
		case '<':
			moveDataPos(-1)
		case '+':
			tape[dataPos]++
		case '-':
			tape[dataPos]--
		case '.':
			fmt.Printf("%c", tape[dataPos])
		case ',':
		case '[':
			if int(datum) == 0 {
				movePastMatchingInstr(1, '[', ']')
			}
		case ']':
			if int(datum) != 0 {
				movePastMatchingInstr(-1, ']', '[')
			}
		default:
			// Skip unrecognized characters
		}
		moveInstrPos(1)
	}
}

func moveDataPos(offset int) {
	dataPos += offset
	if dataPos < 0 || dataPos >= len(tape) {
		log.Fatal("Data pointer out of bounds.")
	}
}

func moveInstrPos(offset int) {
	instrPos += offset
	if instrPos < 0 {
		log.Fatal("Instruction pointer out of bounds.")

	}
	if instrPos >= len(instructions) {
		os.Exit(0)
	}
}

func movePastMatchingInstr(offset int, pair0, pair1 byte) {
	count := 0
	for {
		moveInstrPos(offset)
		instr := instructions[instrPos]
		if instr == pair1 && count == 0 {
			// found
			break
		}
		if instr == pair0 {
			count++
		} else if instr == pair1 {
			count--
			if count < 0 {
				panic("Unreached.")
			}
		}
	}
}
