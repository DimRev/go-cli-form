package demo

import (
	"fmt"
	"time"

	"github.com/DimRev/go-cli-form/internal/form"
)

func RenderDemo() {

	cleanCh, doneCh := form.Clean()
	renderCh, swapCh := form.Render()

	for j := 0; j < 5; j++ {
		l := form.Line{}
		for i := 0; i < 10; i++ {
			colors := []string{"red", "yellow", "blue", "green", "blue", "magenta", "cyan", "white", "grey"}
			l.Part = append(l.Part, form.Part{
				Color:   colors[(i+j)%9],
				Content: fmt.Sprintf("%v,%v", i, j),
			})
		}
		renderCh <- l           // Send the Line object through rCh
		time.Sleep(time.Second) // Simulate some processing time
		swapCh <- j%2 == 0      // Send a cleanup signal after processing
	}
	close(renderCh)
	close(swapCh)
	time.Sleep(time.Second)

	cleanCh <- 1
	<-doneCh
	cleanCh <- 1
	<-doneCh

	close(doneCh)
	close(cleanCh)
}
