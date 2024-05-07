package form

import (
	"fmt"
	"strings"
)

func (f *Form) textStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.TextColor, str))
}

func (f *Form) markedStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.MarkColor, str, f.TextColor))
}

func (f *Form) selectedStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.SelectColor, str, f.TextColor))
}

func redStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["red_fg"], str, ANSI["default_fg"]))
}
func greenStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["green_fg"], str, ANSI["default_fg"]))
}
func yellowStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["yellow_fg"], str, ANSI["default_fg"]))
}
func blueStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["blue_fg"], str, ANSI["default_fg"]))
}
func magentaStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["magenta_fg"], str, ANSI["default_fg"]))
}
func cyanStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["cyan_fg"], str, ANSI["default_fg"])
}
func whiteStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["white_fg"], str, ANSI["default_fg"])
}
func grayStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["gray_fg"], str, ANSI["default_fg"])
}
