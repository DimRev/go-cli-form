package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func (f *Form) SelectInput(qst string, opts []string) (string, error) {
	answerIdx := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(f.SelectedLineIndicator, f.textStrColor(qst, ": "), f.selectOptions(opts, answerIdx, false), "\n")

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

		fmt.Print(MACROS["clean_and_up"], f.SelectedLineIndicator, f.textStrColor(qst, ": "), f.selectOptions(opts, answerIdx, false), "\n")
	}
	fmt.Print(MACROS["clean_and_up"], f.FormTheme.TextColor, f.textStrColor(qst, ": "), f.selectOptions(opts, answerIdx, true), "\n")
	return opts[answerIdx], nil
}

func (f *Form) selectOptions(opts []string, ansIdx int, isLocked bool) string {
	formattedString := ""
	for idx, option := range opts {
		if idx == ansIdx {
			if isLocked {
				formattedString += fmt.Sprint(f.selectedStrColor(option))
			} else {
				formattedString += fmt.Sprint(f.markedStrColor(option))
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
