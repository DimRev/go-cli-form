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

	l1 := Line{BreakLine: true}
	l1.Part = append(l1.Part, Part{
		Color:   f.IndicatorColor,
		Content: f.Indicator,
	}, Part{
		Color:   f.TextColor,
		Content: qst,
	})

	f.renderLineCh <- l1
	f.multiselectOptions(opts, selectedIdx, answerBools, false, false)
	fmt.Print("test 123123")
	f.swapRenderedLineCh <- false
	// fmt.Print(f.Indicator, f.textStrColor(qst, ": "), "\n",
	// 	grayStrColor("(Press space to select option, enter to lock in your selections)"),
	// 	f.multiselectOptions(opts, selectedIdx, answerBools, false, false), "\n",
	// )

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
		if key == keyboard.KeyEsc || key == keyboard.KeyEnter || key == keyboard.KeyCtrlC {
			break
		}

		// fmt.Print(
		// 	MACROS["clean_and_up"], f.Indicator, f.textStrColor(qst, ": "), "\n",
		// 	grayStrColor("(Press space to select option, enter to lock in your selections)"),
		// 	f.multiselectOptions(opts, selectedIdx, answerBools, false, true), "\n",
		// )
		f.multiselectOptions(opts, selectedIdx, answerBools, false, true)
		f.swapRenderedLineCh <- false
	}
	// fmt.Print(MACROS["clean_and_up"], f.textStrColor(qst, ": "), "\n",
	// 	grayStrColor("(Press space to select option, enter to lock in your selections)"),
	// 	f.multiselectOptions(opts, selectedIdx, answerBools, true, true), "\n",
	// )
	f.multiselectOptions(opts, selectedIdx, answerBools, true, true)
	f.swapRenderedLineCh <- false

	return multiselectResults(opts, answerBools), nil
}

func (f *Form) multiselectOptions(opts []string, sel int, ansBools []bool, locked bool, rm bool) {
	if rm {
		f.numLinesCleanCh <- len(opts)
		f.doneCleanCh <- struct{}{}
		// for i := 0; i < len(opts)+1; i++ {
		// fmt.Print(MACROS["clean_and_up"])
		// }
	}

	// formattedString := "\n"
	for idx, option := range opts {
		l1 := Line{BreakLine: false}
		if ansBools[idx] {
			if idx == sel {
				if locked {
					l1.Part = append(l1.Part, Part{
						Color:   f.TextColor,
						Content: f.BulletPointDefault,
					}, Part{
						Color:   f.TextColor,
						Content: option,
					})
					// formattedString += fmt.Sprint(f.BulletPointDefault, f.selectedStrColor(option))
				} else {
					l1.Part = append(l1.Part, Part{
						Color:   f.TextColor,
						Content: f.BulletPointMarked,
					}, Part{
						Color:   f.MarkColor,
						Content: option,
					})
					// formattedString += fmt.Sprint(f.BulletPointMarked, f.markedStrColor(option))
				}
			} else {
				l1.Part = append(l1.Part, Part{
					Color:   f.TextColor,
					Content: f.BulletPointSelected,
				}, Part{
					Color:   f.SelectColor,
					Content: option,
				})
				// formattedString += fmt.Sprint(f.BulletPointSelected, f.selectedStrColor(option))
			}
		} else {
			if idx == sel {
				if locked {
					l1.Part = append(l1.Part, Part{
						Color:   f.TextColor,
						Content: f.BulletPointDefault,
					}, Part{
						Color:   f.TextColor,
						Content: option,
					})
					// formattedString += fmt.Sprint(f.BulletPointDefault, f.textStrColor(option))
				} else {
					l1.Part = append(l1.Part, Part{
						Color:   f.TextColor,
						Content: f.BulletPointDefault,
					}, Part{
						Color:   f.MarkColor,
						Content: option,
					})
					// formattedString += fmt.Sprint(f.BulletPointDefault, f.markedStrColor(option))
				}
			} else {
				l1.Part = append(l1.Part, Part{
					Color:   f.TextColor,
					Content: f.BulletPointDefault,
				}, Part{
					Color:   f.TextColor,
					Content: option,
				})
				// formattedString += fmt.Sprint(f.BulletPointDefault, f.textStrColor(option))
			}
			// f.renderLineCh <- l1
			// f.swapRenderedLineCh <- false
			fmt.Print(l1)
		}

		// if idx < len(opts)-1 {
		// 	formattedString += "\n"
		// }
	}
	// return formattedString
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
