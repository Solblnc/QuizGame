// Quiz Game
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Determine a csv file with flag package
	csvfilename := flag.String("csv", "Problems.csv", "a csv file in format of 'question, answer'")
	// Setting a time limit for a timer
	timelimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// Opening a file
	file, err := os.Open(*csvfilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open a csv file: %s\n", *csvfilename))

	}
	// Reading a csv file
	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse a csv file")
	}
	problems := parselines(lines)

	// Timer
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	// Output of a programme that counts a correct answers
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerch <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerch:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

// Function for parsing questions and answers and implement them to a slice
func parselines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

// Struct: q = questions, a = answers
type problem struct {
	q string
	a string
}

// Exit function for output
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
