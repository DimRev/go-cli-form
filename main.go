package main

import (
	"github.com/DimRev/go-cli-form/form"
)

var ANSI map[string]string
var GREEN_ARROW string
var LINEUP_CLEAR_LINE string

func main() {

	question_1 := "This is the first question"
	question_2 := "This is the second question"
	question_3 := "This is the third question"
	question_4 := "This is the fourth question"
	options1 := []string{"1", "2", "3", "4"}
	options2 := []string{"yes", "no"}
	options3 := []string{"hi", "121", "3", "cool", "5", "hello world", "7"}

	Form := form.Start()
	Form.TextInput(question_1)
	Form.SelectInput(question_2, options1)
	Form.MultiSelectInput(question_3, options2)
	Form.MultiSelectInput(question_4, options3)
	Form.End()
}
