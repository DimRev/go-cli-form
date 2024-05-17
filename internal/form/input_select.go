package form

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func (f *Form) SelectInput(qst string, opts []string) (string, error) {
	answerIdx := 0

	fmt.Print(ANSI["disable_cur"])

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
	l1.Part = append(l1.Part, f.selectOptions(opts, answerIdx, false)...)

	f.renderLineCh <- l1
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
		if key == keyboard.KeyEnter || key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}
		f.swapRenderedLineCh <- true

		l1 = Line{BreakLine: true}
		l1.Part = append(l1.Part, Part{
			Color:   f.IndicatorColor,
			Content: f.Indicator,
		}, Part{
			Color:   f.TextColor,
			Content: qst,
		})
		l1.Part = append(l1.Part, f.selectOptions(opts, answerIdx, false)...)

		f.renderLineCh <- l1
	}

	f.swapRenderedLineCh <- true

	l1 = Line{BreakLine: true}
	l1.Part = append(l1.Part, Part{
		Color:   f.TextColor,
		Content: qst,
	})
	l1.Part = append(l1.Part, f.selectOptions(opts, answerIdx, true)...)

	f.renderLineCh <- l1

	return opts[answerIdx], nil
}

func (f *Form) selectOptions(opts []string, ansIdx int, isLocked bool) []Part {
	answerParts := []Part{}
	for idx, option := range opts {
		if idx == ansIdx {
			if isLocked {
				answerParts = append(answerParts, Part{
					Color:   f.SelectColor,
					Content: option,
				})
			} else {
				answerParts = append(answerParts, Part{
					Color:   f.MarkColor,
					Content: option,
				})
			}
		} else {
			answerParts = append(answerParts, Part{
				Color:   f.TextColor,
				Content: option,
			})
		}
	}
	return answerParts
}
