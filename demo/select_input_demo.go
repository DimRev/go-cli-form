package demo

import (
	"github.com/DimRev/go-cli-form/internal/form"
)

func SelectInputDemo() {
	Form := form.Start("default")
	_, err := Form.SelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	Form.End()
}
