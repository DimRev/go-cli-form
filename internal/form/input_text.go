package form

import "fmt"

func (f *Form) TextInput(qst string) string {
	var answer *string
	fmt.Print(ANSI["enable_cur"])
	fmt.Print(f.SelectedLineIndicator, f.textStrColor(qst), ": ")

	f.Scanner.Scan()
	input := f.Scanner.Text()

	answer = &input

	if *answer != "" {
		fmt.Print(MACROS["clean_and_up"], f.textStrColor(qst, ": "), f.selectedStrColor(*answer), "\n")
	} else {
		fmt.Print(MACROS["clean_and_up"], f.textStrColor(qst, ": "), "\n")
	}

	fmt.Print(ANSI["disable_cur"])
	return *answer
}
