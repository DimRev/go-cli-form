package form

import (
	"fmt"
	"strings"
)

func (f *Form) TextStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.TextColor, str))
}

func (f *Form) MarkedStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.MarkColor, str, f.TextColor))
}

func (f *Form) SelectedStrColor(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(f.SelectColor, str, f.TextColor))
}

func RedStr(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["red_fg"], str, ANSI["default_fg"]))
}
func GreenStr(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["green_fg"], str, ANSI["default_fg"]))
}
func YellowStr(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["yellow_fg"], str, ANSI["default_fg"]))
}
func BlueStr(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["blue_fg"], str, ANSI["default_fg"]))
}
func MagentaStr(s ...string) string {
	str := strings.Join(s, " ")
	return (fmt.Sprint(ANSI["magenta_fg"], str, ANSI["default_fg"]))
}
func CyanStr(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["cyan_fg"], str, ANSI["default_fg"])
}
func WhiteStr(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["white_fg"], str, ANSI["default_fg"])
}
func GrayStr(s ...string) string {
	str := strings.Join(s, " ")
	return fmt.Sprint(ANSI["gray_fg"], str, ANSI["default_fg"])
}
