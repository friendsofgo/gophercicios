package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problem.csv", "path from a quiz file")
	flag.Parse()

	problems, err := parseQuiz(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}

	var correct int
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var userAnswer string
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == problem.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseQuiz(filename string) ([]problem, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open the csv file: %s | %w", filename, err)
	}
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse the provided CSV file | %w", err)
	}
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems, nil
}
