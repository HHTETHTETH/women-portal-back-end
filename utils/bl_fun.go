package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Horodate -
func Horodate(xdate string) int {
	words := strings.Fields(xdate)
	//fmt.Println(words, len(words))

	day, err := strconv.Atoi(words[0])
	if err != nil {
		return 0
	}
	fmt.Println(day)
	month := words[1]
	switch month {
	case "May":
		if day < 22 {
			return 1
		}
		break
	case "April":
		if day < 21 {
			return 0
		}
		break
	case "Jun":
		if day < 21 {
			return 2
		}
		break
	case "Jul":
		if day < 23 {
			return 3
		}
	case "Aug":
		if day < 24 {
			return 4
		}
	case "Sep":
		if day < 23 {
			return 5
		}
		break
	case "Oct":
		if day < 24 {
			return 6
		}
		break
	case "Nov":
		if day < 24 {
			return 7
		}
		break
	case "Dec":
		if day < 22 {
			return 8
		}
		break
	case "Jan":
		if day < 21 {
			return 9
		}
		break
	case "Feb":
		if day < 19 {
			return 10
		}
		break
	case "Mar":
		if day < 21 {
			return 10
		}
		break
	}

	return 0
}
