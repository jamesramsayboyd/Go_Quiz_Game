package main

import (
	"encoding/csv"
	"flag"
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


func loadQuestions(fileName string) []question{

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: File not found.")
		//panic(err)
	}
	//defer file.Close();

	csvContents, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error: Could not read questions.")
		//panic(err)
	}

	var quiz []question


	for _, line := range csvContents {
		qa := question{
			question: line[0],
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

	//fmt.Println("Welcome to the CSV Quiz Game designed in Go Lang");
	fileName := flag.String("questions","problems.csv","The CSV file from where questions would be read.")
	timeLimit := flag.Int("time", 30, "Set a time limit for the quiz")
	flag.Parse()

	var quiz []question
	var score int
	score = 0
	quiz = loadQuestions(*fileName)
	var answer string

	fmt.Printf("Press Enter to start the quiz")
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nTime's up. You answered %d questions correctly out of a total of %d questions.\n",score, len(quiz))
		os.Exit(1)
		}()

	for i, questions := range quiz {
		fmt.Printf("Question %d :- %s? ",i+1,questions.question)
		fmt.Scan(&answer)
		if strings.ToLower(strings.Trim(answer," ")) == strings.ToLower(strings.Trim(questions.answer, " ")){
			score++;
		}
	}

	fmt.Printf("You answered %d questions correctly out of a total of %d questions.\n",score, len(quiz))

}