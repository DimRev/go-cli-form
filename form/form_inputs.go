package form

import (
	"bufio"
	"fmt"

	"github.com/eiannone/keyboard"
)

func FormOpenQuestion(s *bufio.Scanner, q string) string {
	var answer *string
	fmt.Print(ANSI["enable_cur"])
	fmt.Print(GREEN_ARROW, q, ": ")

	s.Scan()
	input := s.Text()

	answer = &input
	if *answer == "" {
		answer = nil
	}

	if answer != nil {
		fmt.Print(LINEUP_CLEAR_LINE, q, ": ", YellowStr(*answer), "\n")
	} else {
		fmt.Print(LINEUP_CLEAR_LINE, q, ": ", "\n")
	}

	fmt.Print(ANSI["disable_cur"])
	return *answer
}

func FormOptionsSelect(q string, o []string) string {
	answerIdx := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(GREEN_ARROW, q, " [", displayOption(o, answerIdx, false), "]\n")

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowLeft:
			if answerIdx == 0 {
				answerIdx = len(o) - 1
			} else {
				answerIdx -= 1
			}

		case keyboard.KeyArrowRight:
			if answerIdx == len(o)-1 {
				answerIdx = 0
			} else {
				answerIdx += 1
			}

		}
		if key == keyboard.KeyEnter || key == keyboard.KeyEsc {
			break
		}

		fmt.Print(LINEUP_CLEAR_LINE, GREEN_ARROW, q, " [", displayOption(o, answerIdx, false), "]\n")
	}
	fmt.Print(LINEUP_CLEAR_LINE, q, " [", displayOption(o, answerIdx, true), "]\n")
	return o[answerIdx]
}

func displayOption(o []string, aIdx int, locked bool) string {
	formattedString := ""
	for idx, option := range o {
		if idx == aIdx {
			if locked {
				formattedString += fmt.Sprint(YellowStr(option))
			} else {
				formattedString += fmt.Sprint(GreenStr(option))
			}
		} else {
			formattedString += fmt.Sprint(option)
		}
		if idx < len(o)-1 {
			formattedString += " / "
		} else {
			formattedString += " "

		}
	}
	return formattedString
}
