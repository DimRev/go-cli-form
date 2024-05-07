package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func (f *Form) MultiSelectInput(qst string, opts []string) ([]string, error) {
	selectedIdx := 0
	answerBools := make([]bool, len(opts))

	err := keyboard.Open()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Print(f.SelectedLineIndicator, f.textStrColor(qst, ": "), "\n",
		grayStrColor("(Press space to select option, enter to lock in your selections)"),
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
			MACROS["clean_and_up"], f.SelectedLineIndicator, f.textStrColor(qst, ": "), "\n",
			grayStrColor("(Press space to select option, enter to lock in your selections)"),
			f.multiselectOptions(opts, selectedIdx, answerBools, false, true), "\n",
		)
	}
	fmt.Print(MACROS["clean_and_up"], f.textStrColor(qst, ": "), "\n",
		grayStrColor("(Press space to select option, enter to lock in your selections)"),
		f.multiselectOptions(opts, selectedIdx, answerBools, true, true), "\n",
	)

	return multiselectResults(opts, answerBools), nil
}

func (f *Form) multiselectOptions(opts []string, sel int, ansBools []bool, locked bool, rm bool) string {
	if rm {
		for i := 0; i < len(opts)+1; i++ {
			fmt.Print(MACROS["clean_and_up"])
		}
	}

	formattedString := "\n"
	for idx, option := range opts {

		if ansBools[idx] {
			if idx == sel {
				if locked {
					formattedString += fmt.Sprint(f.BulletPointDefault, f.selectedStrColor(option))
				} else {
					formattedString += fmt.Sprint(f.BulletPointMarked, f.markedStrColor(option))
				}
			} else {
				formattedString += fmt.Sprint(f.BulletPointSelected, f.selectedStrColor(option))
			}
		} else {
			if idx == sel {
				if locked {
					formattedString += fmt.Sprint(f.BulletPointDefault, f.textStrColor(option))
				} else {
					formattedString += fmt.Sprint(f.BulletPointDefault, f.markedStrColor(option))
				}
			} else {
				formattedString += fmt.Sprint(f.BulletPointDefault, f.textStrColor(option))
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
