package form

import (
	"bufio"
	"fmt"

	"github.com/eiannone/keyboard"
)

func TextInput(s *bufio.Scanner, qst string) string {
	var answer *string
	fmt.Print(ANSI["enable_cur"])
	fmt.Print(GREEN_ARROW, qst, ": ")

	s.Scan()
	input := s.Text()

	answer = &input

	if *answer != "" {
		fmt.Print(LINEUP_CLEAR_LINE, qst, ": ", YellowStr(*answer), "\n")
	} else {
		fmt.Print(LINEUP_CLEAR_LINE, qst, ": ", "\n")
	}

	fmt.Print(ANSI["disable_cur"])
	return *answer
}

func SelectInput(qst string, opts []string) string {
	answerIdx := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(GREEN_ARROW, qst,
		": ", selectOptions(opts, answerIdx, false), "\n",
	)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowLeft:
			if answerIdx == 0 {
				answerIdx = len(opts) - 1
			} else {
				answerIdx -= 1
			}

		case keyboard.KeyArrowRight:
			if answerIdx == len(opts)-1 {
				answerIdx = 0
			} else {
				answerIdx += 1
			}

		}
		if key == keyboard.KeyEnter || key == keyboard.KeyEsc {
			break
		}

		fmt.Print(LINEUP_CLEAR_LINE, GREEN_ARROW, qst,
			": ", selectOptions(opts, answerIdx, false), "\n",
		)
	}
	fmt.Print(LINEUP_CLEAR_LINE, qst,
		": ", selectOptions(opts, answerIdx, true), "\n",
	)
	return opts[answerIdx]
}

func selectOptions(opts []string, ansIdx int, isLocked bool) string {
	formattedString := ""
	for idx, option := range opts {
		if idx == ansIdx {
			if isLocked {
				formattedString += fmt.Sprint(YellowStr(option))
			} else {
				formattedString += fmt.Sprint(GreenStr(option))
			}
		} else {
			formattedString += fmt.Sprint(option)
		}
		if idx < len(opts)-1 {
			formattedString += " / "
		} else {
			formattedString += " "

		}
	}
	return formattedString
}

func MultiSelectInput(qst string, opts []string) []string {
	selectedIdx := 0
	answerBools := make([]bool, len(opts))

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(GREEN_ARROW, qst, ": \n",
		GrayStr("(Press space to select option, enter to lock in your selections)"),
		multiselectOptions(opts, selectedIdx, answerBools, false, false), "\n",
	)

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedIdx == 0 {
				selectedIdx = len(opts) - 1
			} else {
				selectedIdx -= 1
			}

		case keyboard.KeyArrowDown:
			if selectedIdx == len(opts)-1 {
				selectedIdx = 0
			} else {
				selectedIdx += 1
			}
		case keyboard.KeySpace:
			answerBools[selectedIdx] = !answerBools[selectedIdx]
		}
		if key == keyboard.KeyEsc || key == keyboard.KeyEnter {
			break
		}

		fmt.Print(
			LINEUP_CLEAR_LINE, GREEN_ARROW, qst, ": \n",
			GrayStr("(Press space to select option, enter to lock in your selections)"),
			multiselectOptions(opts, selectedIdx, answerBools, false, true), "\n",
		)
	}
	fmt.Print(LINEUP_CLEAR_LINE, qst, ": \n",
		GrayStr("(Press space to select option, enter to lock in your selections)"),
		multiselectOptions(opts, selectedIdx, answerBools, true, true), "\n",
	)

	return multiselectResults(opts, answerBools)
}

func multiselectOptions(opts []string, sel int, ansBools []bool, locked bool, rm bool) string {
	if rm {
		for i := 0; i < len(opts)+1; i++ {
			fmt.Print(LINEUP_CLEAR_LINE)
		}
	}

	formattedString := "\n"
	for idx, option := range opts {

		if ansBools[idx] {
			if idx == sel {
				if locked {
					formattedString += fmt.Sprint(" 🟡  ", YellowStr(option))
				} else {
					formattedString += fmt.Sprint(" 🟡  ", GreenStr(option))
				}
			} else {
				formattedString += fmt.Sprint(" 🟡  ", YellowStr(option))
			}
		} else {
			if idx == sel {
				formattedString += fmt.Sprint(" ⚪  ", GreenStr(option))
			} else {
				formattedString += fmt.Sprint(" ⚪  ", option)
			}
		}

		if idx < len(opts)-1 {
			formattedString += "\n"
		}
	}
	return formattedString
}

func multiselectResults(opts []string, ansBools []bool) []string {
	var res []string
	for idx := range opts {
		if ansBools[idx] {
			res = append(res, opts[idx])
		}
	}
	return res
}
