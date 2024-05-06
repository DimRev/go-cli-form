package form

import (
	"bufio"
	"fmt"
	"os"
)

var ANSI map[string]string
var SYMBOLS map[string]string
var MACROS map[string]string
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

	SYMBOLS = map[string]string{
		"circle_yellow":      " 🟡  ",
		"circle_white":       " ⚪  ",
		"circle_black":       " ⚫  ",
		"circle_red":         " 🔴  ",
		"circle_blue":        " 🔵  ",
		"circle_orange":      " 🟠  ",
		"circle_lightorange": " 🟡  ",
		"circle_green":       " 🟢  ",
		"circle_purple":      " 🟣  ",
		"circle_brown":       " 🟤  ",

		"arrowIndicator": " ➜  ",
	}

	MACROS = map[string]string{
		"clean_and_up": fmt.Sprint(ANSI["cursor_up"], ANSI["delete_line"]),
	}
}

type Form struct {
	Scanner *bufio.Scanner
	FormTheme
}

type FormTheme struct {
	TextColor   string
	MarkColor   string
	SelectColor string

	BulletPointDefault  string
	BulletPointMarked   string
	BulletPointSelected string

	SelectedLineIndicator string
}

type FormInterface interface {
	TextInput(qst string) string
	SelectInput(qst string, opts []string) string
	MultiSelectInput(qst string, opts []string) []string
	End()
}

func NewForm(scanner *bufio.Scanner, theme string) *Form {
	selectedTheme := FormTheme{
		TextColor:   ANSI["default_fg"],
		MarkColor:   ANSI["green_fg"],
		SelectColor: ANSI["yellow_fg"],

		BulletPointDefault:  SYMBOLS["circle_white"],
		BulletPointMarked:   SYMBOLS["circle_green"],
		BulletPointSelected: SYMBOLS["circle_yellow"],

		SelectedLineIndicator: GreenStr(SYMBOLS["arrowIndicator"]),
	}

	if theme == "blue" {
		selectedTheme = FormTheme{
			TextColor:   ANSI["blue_fg"],
			MarkColor:   ANSI["green_fg"],
			SelectColor: ANSI["yellow_fg"],

			BulletPointDefault:  SYMBOLS["circle_blue"],
			BulletPointMarked:   SYMBOLS["circle_green"],
			BulletPointSelected: SYMBOLS["circle_yellow"],

			SelectedLineIndicator: RedStr(SYMBOLS["arrowIndicator"]),
		}
	}

	if theme == "red" {
		selectedTheme = FormTheme{
			TextColor:   ANSI["red_fg"],
			MarkColor:   ANSI["blue_fg"],
			SelectColor: ANSI["green_fg"],

			BulletPointDefault:  SYMBOLS["circle_red"],
			BulletPointMarked:   SYMBOLS["circle_blue"],
			BulletPointSelected: SYMBOLS["circle_green"],

			SelectedLineIndicator: YellowStr(SYMBOLS["arrowIndicator"]),
		}
	}

	return &Form{
		Scanner:   scanner,
		FormTheme: selectedTheme,
	}
}

func Start(theme string) Form {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(ANSI["disable_cur"])
	return *NewForm(scanner, theme)
}

func (f *Form) End() {
	fmt.Print(ANSI["enable_cur"])
}
