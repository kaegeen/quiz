package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func main() {
	questions := []Question{
		{"What is the capital of France?", "Paris"},
		{"What is 5 + 7?", "12"},
		{"What is the color of the sky on a clear day?", "Blue"},
		{"Who wrote 'To Kill a Mockingbird'?", "Harper Lee"},
		{"What is the largest planet in our solar system?", "Jupiter"},
	}

	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("==========================")
	fmt.Println("Instructions:")
	fmt.Println(" - Answer each question to the best of your ability.")
	fmt.Println(" - Your answers are case-insensitive.")
	fmt.Println(" - Type 'exit' to quit the game early.")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	score := 0

	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.question)
		fmt.Print("Your answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.ToLower(answer) == "exit" {
			fmt.Println("You chose to exit the quiz.")
			break
		}

		if strings.EqualFold(answer, q.answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is: %s\n", q.answer)
		}
	}

	fmt.Printf("\nQuiz finished! Your final score is: %d out of %d\n", score, len(questions))
	fmt.Println("Thanks for playing!")
}
