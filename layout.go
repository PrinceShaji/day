package day

import "strings"

var layoutReplacements = []string{
	"YYYY", "2006",
	"YYY", "006", // new
	"YY", "06",
	"MMMM", "January",
	"MMM", "Jan",
	"MM", "01",
	"M", "1",
	"DD", "02",
	"D", "2",
	"dddd", "MOnday", // new
	"ddd", "Mon", // new
	// "dd", "nil", // Mo,  not available
	// "d", "nil", // 0-6,  not available
	"HH", "15", // 05 zero padded 24 hr, curently not available
	"H", "15",
	"hh", "03",
	"h", "3",
	"mm", "04",
	"m", "4",
	"SSSS", "124006", // nanosecond
	"ss", "05",
	"s", "5",

	// AM/PM
	"A", "PM",
	"a", "pm",

	// Timezones
	"ZZ", "-07", // Timezone offset without colon  eg: -0200
	"Z", "-07:00", // Timezone offset with colon. eg: -07:00
}

func generateLayout(layout string) string {
	r := strings.NewReplacer(layoutReplacements...)

	return r.Replace(layout)
}
