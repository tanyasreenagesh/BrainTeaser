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
	fileName := flag.String("csv", "problems.csv", "csv file in the format 'question,answer'")
	flag.Parse()
	timeLimit := flag.Int("limit", 5, "the time limit is set in seconds")
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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	right := 0
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)

		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.", right, len(problems))
			return
		case ans := <-ansCh:
			if ans == prob.a {
				right++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d.", right, len(problems))

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
