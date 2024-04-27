package main

import (
	"bufio"
	"fmt"
	"os"
)

var ANSI map[string]string

func main() {
	ANSI = map[string]string{
		"cursor_up":   "\x1b[1A",
		"delete_line": "\x1b[2K",
	}

	fmt.Println("Hello world")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Starting here:")

	question_1 := "This is the first question"

	_ = formOpenQuestion(scanner, question_1)

}

func formOpenQuestion(s *bufio.Scanner, q string) string {
	var answer *string

	fmt.Print("➜ ")
	fmt.Print(q, ": ")

	s.Scan()
	input := s.Text()

	answer = &input
	if *answer == "" {
		answer = nil
	}

	if answer != nil {
		fmt.Print(ANSI["cursor_up"], ANSI["delete_line"], q, ": ", *answer, "\n")
	} else {
		fmt.Print(ANSI["cursor_up"], ANSI["delete_line"], q, ": ", "\n")
	}

	return *answer
}
