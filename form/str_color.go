package form

import "fmt"

func RedStr(s string) string {
	return fmt.Sprint(ANSI["red_fg"], s, ANSI["default_fg"])
}
func GreenStr(s string) string {
	return fmt.Sprint(ANSI["green_fg"], s, ANSI["default_fg"])
}
func YellowStr(s string) string {
	return fmt.Sprint(ANSI["yellow_fg"], s, ANSI["default_fg"])
}
func BlueStr(s string) string {
	return fmt.Sprint(ANSI["blue_fg"], s, ANSI["default_fg"])
}
func MagentaStr(s string) string {
	return fmt.Sprint(ANSI["magenta_fg"], s, ANSI["default_fg"])
}
func CyanStr(s string) string {
	return fmt.Sprint(ANSI["cyan_fg"], s, ANSI["default_fg"])
}
func WhiteStr(s string) string {
	return fmt.Sprint(ANSI["white_fg"], s, ANSI["default_fg"])
}
func GrayStr(s string) string {
	return fmt.Sprint(ANSI["gray_fg"], s, ANSI["default_fg"])
}