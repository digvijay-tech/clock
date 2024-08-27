// Current Local Time
package clt

import (
	"fmt"
	"time"
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

	formattedTime := currentTime.Format("Mon 2") + suffix + currentTime.Format(" Jan 2006, 03:04 PM MST")
	fmt.Println("Current Local Time:", formattedTime)
}
