package form

import (
	"errors"
	"fmt"
)

func (f *Form) TextInput(qst string) (string, error) {
	if qst == "" {
		return "", errors.New("question cannot be empty")
	}
	var answer *string
	fmt.Print(ANSI["enable_cur"])
	l1 := Line{
		BreakLine: false,
	}
	l1.Part = append(l1.Part, Part{
		Color:   f.IndicatorColor,
		Content: f.Indicator,
	}, Part{
		Color:   f.TextColor,
		Content: qst,
	})
	f.renderLineCh <- l1

	f.Scanner.Scan()
	input := f.Scanner.Text()

	answer = &input

	f.swapRenderedLineCh <- true
	l2 := Line{
		BreakLine: true,
	}
	if *answer != "" {
		l2.Part = append(l2.Part, Part{
			Color:   f.TextColor,
			Content: qst,
		}, Part{
			Color:   f.SelectColor,
			Content: *answer,
		})
	} else {
		l2.Part = append(l2.Part, Part{
			Color:   f.TextColor,
			Content: qst,
		})
	}

	f.renderLineCh <- l2
	f.swapRenderedLineCh <- false

	fmt.Print(ANSI["disable_cur"])
	return *answer, nil
}
