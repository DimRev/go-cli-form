package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DimRev/go-cli-form/form"
)

var ANSI map[string]string
var GREEN_ARROW string
var LINEUP_CLEAR_LINE string

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	question_1 := "This is the first question"
	question_2 := "This is the second question"
	question_3 := "This is the third question"
	question_4 := "This is the fourth question"
	options1 := []string{"1", "2", "3", "4"}
	options2 := []string{"yes", "no"}
	options3 := []string{"hi", "121", "3", "cool", "5", "hello world", "7"}

	form.Start()
	form.TextInput(scanner, question_1)
	form.SelectInput(question_2, options1)
	form.SelectInput(question_3, options2)
	res := form.MultiSelectInput(question_4, options3)
	fmt.Println(res)
	form.End()
}
