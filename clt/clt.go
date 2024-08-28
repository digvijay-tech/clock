// Current Local Time
package clt

import (
	"fmt"
	"time"
)

const (
	DAYALPHA = "Mon"
	DAYNUM   = "2"
	MONTH    = "Jan"
	YEAR     = "2006"
	DATE     = "2 Jan 2006"
	HOUR12   = "03:04 PM"
	HOUR24   = "15:04"
	TIMEZONE = "MST"
)

func getSuffix(day int) string {
	switch day {
	case 1, 21, 31:
		return "st"
	case 2, 22:
		return "nd"
	case 3, 23:
		return "rd"
	default:
		return "th"
	}
}

func PrintCLT() {
	currentTime := time.Now()

	day := currentTime.Day()
	suffix := getSuffix(day)

	formattedTime := currentTime.Format(DAYALPHA+" "+DAYNUM) + suffix + currentTime.Format(" "+MONTH+" "+YEAR+" "+HOUR12+" "+TIMEZONE)
	fmt.Println(formattedTime)
}
