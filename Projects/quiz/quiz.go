package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("timeLimit", 30, "time limit for the quiz in seconds")
	flag.Parse()

	f, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
	}

	r := csv.NewReader(f)
	var correct int
	var lineCounter int

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for {

		record, err := r.Read()

		if err == io.EOF {
			fmt.Printf("\nYou scored %d out of %d", correct, lineCounter)
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		p := parseRecord(record)
		lineCounter++

		fmt.Printf("\nProblem: %v = ", p.q)

		answerCh := make(chan string)
		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d", correct, lineCounter)
			return
		case input := <-answerCh:
			if input == p.a {
				correct++
			}
		}
	}
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
