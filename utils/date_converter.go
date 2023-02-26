package utils

import "time"

// função que converte data para string
func DateToString(date time.Time) string {
	return date.Format("02/01/2006")
}

// função que converte string para data
func StringToDate(date string) time.Time {
	layout := "02/01/2006"
	result, _ := time.Parse(layout, date)
	return result
}
