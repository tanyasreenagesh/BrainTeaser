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
	timeLimit := flag.Int("limit", 60, "the default time limit is 60 seconds")
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
		printProblem(i, prob)

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
	fmt.Printf("\nYou scored %d out of %d.\n", right, len(problems))
	fmt.Println("")
}

func printProblem(i int, prob problem) {
	fmt.Printf("\nQUESTION %d", i+1)
	fmt.Printf("\n------------")
	fmt.Printf("\n%s\n\n", prob.q)
	fmt.Printf("A. %s\nB. %s\nC. %s\nD. %s\n\n", prob.opt1, prob.opt2, prob.opt3, prob.opt4)
	fmt.Printf("Answer: ")
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q:    line[0],
			opt1: line[1],
			opt2: line[2],
			opt3: line[3],
			opt4: line[4],
			a:    strings.TrimSpace(line[5]),
		}
	}

	return ret
}

type problem struct {
	q    string
	opt1 string
	opt2 string
	opt3 string
	opt4 string
	a    string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
