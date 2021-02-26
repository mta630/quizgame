package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	csvGenerator()
	// numberCorrect := 0

	// Open the File
	quizFile, err := os.Open("./quiz.csv")
	if err != nil {
		fmt.Println("An error occured:: ", err)
	}

	// Initialize the reader
	reader := csv.NewReader(quizFile)

	// Read the records and store in variable
	quiz, _ := reader.ReadAll()

	fmt.Println(quiz)

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
