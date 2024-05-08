package form

func defaultTheme() FormTheme {
	return FormTheme{
		TextColor:   "default",
		MarkColor:   "green",
		SelectColor: "yellow",

		BulletPointDefault:  SYMBOLS["circle_white"],
		BulletPointMarked:   SYMBOLS["circle_green"],
		BulletPointSelected: SYMBOLS["circle_yellow"],

		IndicatorColor: "green",
		Indicator:      SYMBOLS["arrowIndicator"],
	}
}

func blueTheme() FormTheme {
	return FormTheme{
		TextColor:   "blue",
		MarkColor:   "green",
		SelectColor: "yellow",

		BulletPointDefault:  SYMBOLS["circle_blue"],
		BulletPointMarked:   SYMBOLS["circle_green"],
		BulletPointSelected: SYMBOLS["circle_yellow"],

		IndicatorColor: "red",
		Indicator:      SYMBOLS["arrowIndicator"],
	}
}

func redTheme() FormTheme {
	return FormTheme{
		TextColor:   "red",
		MarkColor:   "blue",
		SelectColor: "green",

		BulletPointDefault:  SYMBOLS["circle_red"],
		BulletPointMarked:   SYMBOLS["circle_blue"],
		BulletPointSelected: SYMBOLS["circle_green"],

		IndicatorColor: "yellow",
		Indicator:      SYMBOLS["arrowIndicator"],
	}
}
