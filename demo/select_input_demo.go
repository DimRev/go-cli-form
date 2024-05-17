package demo

import (
	"github.com/DimRev/go-cli-form/internal/form"
)

func SelectInputDemo() {
	f1 := form.Start("default")
	_, err := f1.SelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	f1.End()
	f2 := form.Start("blue")
	_, err = f2.SelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	f2.End()
	f3 := form.Start("red")
	_, err = f3.SelectInput("Test question", []string{"1", "2", "3"})
	if err != nil {
		println(err)
		panic(1)
	}
	f3.End()
}
