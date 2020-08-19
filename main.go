package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "csv file in the format 'question,answer'")
	flag.Parse()
	file, err := os.Open(*fileName)

	if err != nil {
		exit(fmt.Sprintf("The file named %s could not be opened! :(", *fileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	problems := parseLines(lines)
	fmt.Println(problems)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
