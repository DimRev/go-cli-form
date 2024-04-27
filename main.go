package main

import (
	"bufio"
	"fmt"
	"os"
)

var ANSI map[string]string
var GREEN_ARROW string
var LINEUP_CLEAR_LINE string

func main() {
	ANSI = map[string]string{
		"cursor_up":   "\x1b[1A",
		"disable_cur": "\x1b[?25l",
		"enable_cur":  "\x1b[?25h",

		"delete_line": "\x1b[2K",

		"black_fg":   "\x1b[30m",
		"red_fg":     "\x1b[31m",
		"green_fg":   "\x1b[32m",
		"yellow_fg":  "\x1b[33m",
		"blue_fg":    "\x1b[34m",
		"magenta_fg": "\x1b[35m",
		"cyan_fg":    "\x1b[36m",
		"white_fg":   "\x1b[37m",
		"default_fg": "\x1b[39m",

		"left_arrow":  "\x1b[D",
		"right_arrow": "\x1b[C",
	}

	GREEN_ARROW = fmt.Sprint(GreenStr("➜ "))
	LINEUP_CLEAR_LINE = fmt.Sprint(ANSI["cursor_up"], ANSI["delete_line"])

	scanner := bufio.NewScanner(os.Stdin)

	question_1 := "This is the first question"
	question_2 := "This is the second question"
	question_3 := "This is the third question"
	question_4 := "This is the fourth question"
	options1 := []string{"1", "2", "3", "4"}
	options2 := []string{"yes", "no"}
	options3 := []string{"cool", "no"}

	fmt.Print(ANSI["disable_cur"])
	_ = FormOpenQuestion(scanner, question_1)
	_ = FormOptionsSelect(question_2, options1)
	_ = FormOptionsSelect(question_3, options2)
	_ = FormOptionsSelect(question_4, options3)
	fmt.Print(ANSI["enable_cur"])
}
