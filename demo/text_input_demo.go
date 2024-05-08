package demo

import (
	"fmt"

	"github.com/DimRev/go-cli-form/internal/form"
)

func TextInputDemo() {
	Form1 := form.Start("default")
	answer1, err := Form1.TextInput("Testing the question1")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	answer2, err := Form1.TextInput("Testing the question2")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	Form1.End()

	Form2 := form.Start("blue")
	answer3, err := Form2.TextInput("Testing the question3")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	answer4, err := Form2.TextInput("Testing the question4")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	Form2.End()

	Form3 := form.Start("red")
	answer5, err := Form3.TextInput("Testing the question3")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	answer6, err := Form3.TextInput("Testing the question4")
	if err != nil {
		fmt.Print(err)
		panic(1)
	}
	Form3.End()

	fmt.Println("answers: ", answer1, answer2, answer3, answer4, answer5, answer6)
}
