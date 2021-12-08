package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
	}

	r := csv.NewReader(f)
	var correct int
	var lineCounter int

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		p := parseRecord(record)
		lineCounter++

		fmt.Printf("\nProblem: %v = ", p.q)
		var input string
		fmt.Scanln(&input)

		if input == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d", correct, lineCounter)
}

func parseRecord(record []string) problem {
	return problem{
		q: record[0],
		a: strings.TrimSpace(record[1]),
	}
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
