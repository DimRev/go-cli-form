package form

import (
	"bufio"
	"fmt"
	"os"
)

var ANSI map[string]string
var GREEN_ARROW string
var LINEUP_CLEAR_LINE string

func init() {
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
		"gray_fg":    "\x1b[90m",
		"default_fg": "\x1b[39m",

		"left_arrow":  "\x1b[D",
		"right_arrow": "\x1b[C",
	}

	GREEN_ARROW = fmt.Sprint(GreenStr(" ➜  "))
	LINEUP_CLEAR_LINE = fmt.Sprint(ANSI["cursor_up"], ANSI["delete_line"])
}

type Form struct {
	Scanner *bufio.Scanner
}

type FormInterface interface {
	TextInput(qst string) string
	SelectInput(qst string, opts []string) string
	MultiSelectInput(qst string, opts []string) []string
	End()
}

func NewForm(scanner *bufio.Scanner) *Form {
	return &Form{Scanner: scanner}
}

func Start() Form {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(ANSI["disable_cur"])
	return *NewForm(scanner)
}

func (f *Form) End() {
	fmt.Print(ANSI["enable_cur"])
}
