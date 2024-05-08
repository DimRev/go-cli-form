package form

import "fmt"

type Part struct {
	Color   string
	Content string
}

type Line struct {
	Part []Part
}

// Render a line to the terminal with a swap ch that blocks the routing waiting
func Render() (lineCh chan Line, toSwapCh chan bool) {
	lineCh = make(chan Line)
	toSwapCh = make(chan bool)
	go func() {
		for l := range lineCh {
			str := ""
			for _, p := range l.Part {
				str += colorizeString(p)
			}
			fmt.Println(str)
			c := <-toSwapCh // Blocks waiting for user input
			if c {
				fmt.Print(MACROS["clean_and_up"])
			}
		}
	}()

	return lineCh, toSwapCh
}

// Cleans a number of lines returning a done ch on finish
func Clean() (numLinesCh chan int, doneCh chan struct{}) {
	doneCh = make(chan struct{})
	numLinesCh = make(chan int)

	go func() {
		for n := range numLinesCh {
			str := ""
			for i := 0; i < n; i++ {
				str += MACROS["clean_and_up"]
			}
			fmt.Print(str)
			doneCh <- struct{}{}
		}
	}()

	return numLinesCh, doneCh
}

func colorizeString(p Part) string {
	switch p.Color {
	case "red":
		return redStrColor(p.Content, " ")
	case "green":
		return greenStrColor(p.Content, " ")
	case "yellow":
		return yellowStrColor(p.Content, " ")
	case "blue":
		return blueStrColor(p.Content, " ")
	case "magenta":
		return magentaStrColor(p.Content, " ")
	case "cyan":
		return cyanStrColor(p.Content, " ")
	case "white":
		return whiteStrColor(p.Content, " ")
	case "gray":
		return grayStrColor(p.Content, " ")
	default:
		return defaultStrColor(p.Content, " ")
	}
}
