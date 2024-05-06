package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func (f *Form) TextInput(qst string) string {
	var answer *string
	fmt.Print(ANSI["enable_cur"])
	fmt.Print(GREEN_ARROW, f.TextStrColor(qst), ": ")

	f.Scanner.Scan()
	input := f.Scanner.Text()

	answer = &input

	if *answer != "" {
		fmt.Print(LINEUP_CLEAR_LINE, f.TextStrColor(qst+": "), f.SelectedStrColor(*answer), "\n")
	} else {
		fmt.Print(LINEUP_CLEAR_LINE, f.TextStrColor(qst+": "), "\n")
	}

	fmt.Print(ANSI["disable_cur"])
	return *answer
}

func (f *Form) SelectInput(qst string, opts []string) string {
	answerIdx := 0

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(GREEN_ARROW, f.TextStrColor(qst+": "), f.selectOptions(opts, answerIdx, false), "\n")

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

		fmt.Print(LINEUP_CLEAR_LINE, GREEN_ARROW, f.TextStrColor(qst+": "), f.selectOptions(opts, answerIdx, false), "\n")
	}
	fmt.Print(LINEUP_CLEAR_LINE, f.FormTheme.TextColor, f.TextStrColor(qst+": "), f.selectOptions(opts, answerIdx, true), "\n")
	return opts[answerIdx]
}

func (f *Form) selectOptions(opts []string, ansIdx int, isLocked bool) string {
	formattedString := ""
	for idx, option := range opts {
		if idx == ansIdx {
			if isLocked {
				formattedString += fmt.Sprint(f.SelectedStrColor(option))
			} else {
				formattedString += fmt.Sprint(f.MarkedStrColor(option))
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

func (f *Form) MultiSelectInput(qst string, opts []string) []string {
	selectedIdx := 0
	answerBools := make([]bool, len(opts))

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(GREEN_ARROW, f.TextStrColor(qst+": "), "\n",
		GrayStr("(Press space to select option, enter to lock in your selections)"),
		f.multiselectOptions(opts, selectedIdx, answerBools, false, false), "\n",
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
			LINEUP_CLEAR_LINE, GREEN_ARROW, f.TextStrColor(qst+": "), "\n",
			GrayStr("(Press space to select option, enter to lock in your selections)"),
			f.multiselectOptions(opts, selectedIdx, answerBools, false, true), "\n",
		)
	}
	fmt.Print(LINEUP_CLEAR_LINE, f.TextStrColor(qst+": "), "\n",
		GrayStr("(Press space to select option, enter to lock in your selections)"),
		f.multiselectOptions(opts, selectedIdx, answerBools, true, true), "\n",
	)

	return multiselectResults(opts, answerBools)
}

func (f *Form) multiselectOptions(opts []string, sel int, ansBools []bool, locked bool, rm bool) string {
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
					formattedString += fmt.Sprint(f.BulletPointDefault, f.SelectedStrColor(option))
				} else {
					formattedString += fmt.Sprint(f.BulletPointMarked, f.MarkedStrColor(option))
				}
			} else {
				formattedString += fmt.Sprint(f.BulletPointSelected, f.SelectedStrColor(option))
			}
		} else {
			if idx == sel {
				if locked {
					formattedString += fmt.Sprint(f.BulletPointDefault, f.TextStrColor(option))
				} else {
					formattedString += fmt.Sprint(f.BulletPointDefault, f.MarkedStrColor(option))
				}
			} else {
				formattedString += fmt.Sprint(f.BulletPointDefault, f.TextStrColor(option))
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
