package title

import "time"

func ParseCreationDate(dateStr string) time.Time {
	timeParsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}
	return timeParsed
}
