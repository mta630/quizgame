package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	numberCorrect := 0

	// Open the File
	quizFile, err := os.Open("./quiz.csv")
	if err != nil {
		fmt.Println("An error occured:: ", err)
	}

	// Initialize the reader
	reader := csv.NewReader(quizFile)

	// Read the records and store in variable
	quiz, _ := reader.ReadAll()

	// Iterate through
	for _, line := range quiz {
		var answer string
		fmt.Println("Please enter your answer for", line[0])
		fmt.Scanln(&answer)

		if answer == line[1] {
			numberCorrect++
			fmt.Println("You are correct!!")
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("You are not correct!")
			time.Sleep(1 * time.Second)
		}
		clearScreen()
	}

	defer fmt.Println("Your final score was", numberCorrect, "out of", len(quiz))

}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func csvGenerator() {
	questionAnswers := [][]string{
		{"5+5", "10"},
		{"7+3", "10"},
		{"1+5", "6"},
		{"45-13", "32"},
		{"2+4", "6"},
		{"5x5", "25"},
		{"5x=25", "5"},
		{"4-3", "1"},
		{"10/2", "5"},
		{"5x6", "30"},
	}

	csvFile, err := os.Create("quiz.csv")

	if err != nil {
		fmt.Println("failed creating file:", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	for _, quizRow := range questionAnswers {
		_ = csvwriter.Write(quizRow)
	}

	csvwriter.Flush()
}
