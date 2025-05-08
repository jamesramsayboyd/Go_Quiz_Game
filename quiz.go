package main

import (
	"encoding/csv"
	//"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Struct representing a question. Standard OOP class is not present in Go
type question struct{
	question string
	answer string
}

// Function that reads questions from CSV file and stores them as a randomised list of question type structs
func loadQuestions(fileName string) []question{

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: File not found.")
	}

	csvContents, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error: Could not read questions.")
	}

	var quiz []question


	for _, line := range csvContents {
		unformattedQuestion := line[0]
		qa := question{
			question: strings.ReplaceAll(unformattedQuestion, `\n`, "\n"),	// Go doesn't parse "\n" as a newline character
			answer: line[1],
		}
		quiz = append(quiz, qa)
	}
	
	// Random number generation to shuffle question order
	r := rand.New(rand.NewSource(time.Now().Unix()))
		for n := len(quiz); n > 0; n-- {
			randIndex := r.Intn(n)
			quiz[n-1], quiz[randIndex] = quiz[randIndex], quiz[n-1]
		}
	
	return quiz
}



func main(){

	var fileName = "problems.csv"
	var timeLimit = 20
	//flag.Parse()

	var quiz []question
	var score int
	score = 0
	quiz = loadQuestions(fileName)
	var answer string

	fmt.Printf("Welcome to the 20 second quiz. Use the number keys 1, 2, 3 to answer. Press Enter to begin")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(timeLimit)*time.Second)
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nTime's up. Your score: %d out of %d\n",score, len(quiz))
		os.Exit(1)
		}()

	for i, questions := range quiz {
		fmt.Printf("\nQuestion %d: %s ",i+1,questions.question)
		fmt.Scan(&answer)
		if strings.ToLower(strings.Trim(answer," ")) == strings.ToLower(strings.Trim(questions.answer, " ")){
			score++;
		}
	}

	fmt.Printf("End of quiz. Your score: %d out of %d\n",score, len(quiz))

}