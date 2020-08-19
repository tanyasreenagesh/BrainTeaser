package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	right := 0
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)
		var ans string
		fmt.Scanf("%s\n", &ans)

		if ans == prob.a {
			right++
		}
	}

	fmt.Printf("You scored %d out of %d", right, len(problems))
	fmt.Println("")

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
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
