package demo

import (
	"fmt"

	"github.com/DimRev/go-cli-form/internal/form"
)

var ANSI map[string]string
var GREEN_ARROW string
var LINEUP_CLEAR_LINE string

func FormDemo() {

	question_1 := "What is your name"
	question_2 := "What programming languages do you know"
	question_3 := "How long have you been coding(years)"
	question_4 := "What types of projects have you done (side projects included)"

	options2 := []string{"JavaScript", "TypeScript", "GoLang", "Python", "Ruby", "Rust", "C", "C++", "C#", "Zig", "Lua"}
	options3 := []string{"< 1", "1", "2", "3", "4", "5+"}
	options4 := []string{"WebApp", "Bot", "Crawler", "CLI-Tool", "Game", "Mesh", "3D-Animation", "Embedded", "Back End Infra"}

	fmt.Println(FORM_LOGO)

	Form := form.Start("red")
	res1, err := Form.TextInput(question_1)
	if err != nil {
		fmt.Println(err)
	}
	res2, err := Form.MultiSelectInput(question_2, options2)
	if err != nil {
		fmt.Println(err)
	}
	res3, err := Form.SelectInput(question_3, options3)
	if err != nil {
		fmt.Println(err)
	}
	res4, err := Form.MultiSelectInput(question_4, options4)
	if err != nil {
		fmt.Println(err)
	}
	Form.End()

	fmt.Print(
		"result1: ", res1, "\n",
		"result2: ", res2, "\n",
		"result3: ", res3, "\n",
		"result4: ", res4, "\n",
	)
}

var FORM_LOGO = fmt.Sprint("<!-- :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: -->\n",
	"<!-- :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: -->\n",
	"<!-- ::                                                                                        :: -->\n",
	"<!-- ::     _________  .____     .___         ___________________   __________    _____        :: -->\n",
	"<!-- ::     \\_   ___ \\ |    |    |   |        \\_   _____/\\_____  \\  \\______   \\  /     \\       :: -->\n",
	"<!-- ::     /    \\  \\/ |    |    |   |         |    __)   /   |   \\  |       _/ /  \\ /  \\      :: -->\n",
	"<!-- ::     \\     \\____|    |___ |   |         |     \\   /    |    \\ |    |   \\/    Y    \\     :: -->\n",
	"<!-- ::      \\______  /|_______ \\|___|         \\___  /   \\_______  / |____|_  /\\____|__  /     :: -->\n",
	"<!-- ::             \\/         \\/                  \\/            \\/         \\/         \\/      :: -->\n",
	"<!-- ::                                                                                        :: -->\n",
	"<!-- :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: -->\n",
	"<!-- :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: -->")
