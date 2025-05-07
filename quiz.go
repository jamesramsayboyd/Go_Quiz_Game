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

type question struct{
	ques string
	ans string
}

func generateQuestions(fileName string, shuffle bool) []question{

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error :- The CSV file could not be read.");
		panic(err)
	}
	defer file.Close();

	quesAns, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error :- The questions from the CSV file could not be read.");
		panic(err)
	}

	var quiz []question


	for _, line := range quesAns {
		qa := question{
			ques: line[0],
			ans: line[1],
		}
		quiz = append(quiz, qa)
	}
	if shuffle {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for n := len(quiz); n > 0; n-- {
			randIndex := r.Intn(n)
			quiz[n-1], quiz[randIndex] = quiz[randIndex], quiz[n-1]
		}
	}
	return quiz
}



func main(){

	fmt.Println("Welcome to the CSV Quiz Game designed in Go Lang");
	fileName := flag.String("questions","problems.csv","The CSV file from where questions would be read.")
	timeLimit := flag.Int("time", 30, "Set a time limit for the quiz")
	shuffle := flag.Bool("shuffle",false,"Shuffle the order of questions")
	flag.Parse()

	var quiz []question
	var score int
	score = 0
	quiz = generateQuestions(*fileName, *shuffle)
	var answer string

	fmt.Printf("Press Enter key to start the quiz. Duration of quiz :- %d seconds\n",*timeLimit)
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\nTime's up. You answered %d questions correctly out of a total of %d questions.\n",score, len(quiz))
		os.Exit(1)
		}()

	for i, questions := range quiz {
		fmt.Printf("Question %d :- %s? ",i+1,questions.ques)
		fmt.Scan(&answer)
		if strings.ToLower(strings.Trim(answer," ")) == strings.ToLower(strings.Trim(questions.ans, " ")){
			score++;
		}
	}

	fmt.Printf("You answered %d questions correctly out of a total of %d questions.\n",score, len(quiz))

}