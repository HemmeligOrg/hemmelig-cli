package core

import "time"

func Year() int {
	year, _, _ := time.Now().Date()

	return year
}
