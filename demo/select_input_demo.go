package demo

import (
	"fmt"

	"github.com/DimRev/go-cli-form/internal/form"
)

func SelectInputDemo() {
	Form := form.Start("default")
	answer, err := Form.SelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	Form.End()

	fmt.Println(answer)
}
