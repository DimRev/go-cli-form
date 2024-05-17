package demo

import "github.com/DimRev/go-cli-form/internal/form"

func MultiselectInputDemo() {
	f1 := form.Start("default")
	_, err := f1.MultiSelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	f1.End()
}
