package form

func defaultTheme() FormTheme {
	return FormTheme{
		TextColor:   ANSI["default_fg"],
		MarkColor:   ANSI["green_fg"],
		SelectColor: ANSI["yellow_fg"],

		BulletPointDefault:  SYMBOLS["circle_white"],
		BulletPointMarked:   SYMBOLS["circle_green"],
		BulletPointSelected: SYMBOLS["circle_yellow"],

		SelectedLineIndicator: greenStrColor(SYMBOLS["arrowIndicator"]),
	}
}

func blueTheme() FormTheme {
	return FormTheme{
		TextColor:   ANSI["blue_fg"],
		MarkColor:   ANSI["green_fg"],
		SelectColor: ANSI["yellow_fg"],

		BulletPointDefault:  SYMBOLS["circle_blue"],
		BulletPointMarked:   SYMBOLS["circle_green"],
		BulletPointSelected: SYMBOLS["circle_yellow"],

		SelectedLineIndicator: redStrColor(SYMBOLS["arrowIndicator"]),
	}
}

func redTheme() FormTheme {
	return FormTheme{
		TextColor:   ANSI["red_fg"],
		MarkColor:   ANSI["blue_fg"],
		SelectColor: ANSI["green_fg"],

		BulletPointDefault:  SYMBOLS["circle_red"],
		BulletPointMarked:   SYMBOLS["circle_blue"],
		BulletPointSelected: SYMBOLS["circle_green"],

		SelectedLineIndicator: yellowStrColor(SYMBOLS["arrowIndicator"]),
	}
}
