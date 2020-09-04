package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	//Flags setup
	csvPath := flag.String("csvPath", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("timeLimit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	records := readCsvFile(*csvPath)
	problems := parseRecords(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	var numCorrect int

	for i, val := range problems {
		fmt.Printf("Problem #%v: %s = ", i+1, val.q)
		answerCh := make(chan string)

		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answerCh <- scanner.Text()

		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.", numCorrect, len(records))
			return
		case answer := <-answerCh:
			if len(answer) > 0 {
				if answer == val.a {
					numCorrect++
				}
			} else {
				exit("Please enter a number greater than 0.")
			}
		}
	}

	fmt.Printf("You scored %d out of %d.", numCorrect, len(records))
}

func readCsvFile(filePath string) (records [][]string) {
	//Retrieve and open CSV file using csvFilePath
	f, e := os.Open(filePath)
	if e != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", filePath))
	}
	//Close file at the end of the function
	defer f.Close()
	//Declare and initialize new CSV reader
	r := csv.NewReader(f)
	//Read each line of CSV file and convert to a slice of string slices
	records, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return records
}

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))

	for i, val := range records {
		ret[i] = problem{
			q: val[0],
			a: strings.TrimSpace(val[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
