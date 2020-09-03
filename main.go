package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var csvPath = flag.String("csvPath", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()
	var numCorrect int
	records := readCsvFile(*csvPath)
	problems := parseRecords(records)

	for i, val := range problems {
		fmt.Printf("Problem #%v: %s = ", i+1, val.q)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		res := scanner.Text()
		if len(res) > 0 {
			if res == val.a {
				numCorrect++
			}
		} else {
			exit("Please enter a number.")
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

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
