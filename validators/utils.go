package validators

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func Matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func ReplacePattern(str, pattern, replace string) string {
	r, _ := regexp.Compile(pattern)
	return r.ReplaceAllString(str, replace)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Contains(str, substring string) bool {
	return strings.Contains(str, substring)
}

func IsDate(date string) (isValid bool, dateTime time.Time) {

	var (
		day       string
		month     string
		year      string
		hour      string
		min       string
		sec       string
		formatted string
	)

	//Separate date and hours
	datetime := strings.Split(date, " ")
	dateOnly := datetime[0]

	if Contains(date, "T") {
		datetime = strings.Split(date, "T")
		dateOnly = datetime[0]
	}

	//Check date is BR valid
	regex := "^([0-9]{1,2}\\/[0-9]{1,2}\\/[0-9]{4})"
	isBr := Matches(date, regex)

	if isBr {
		dt := strings.Split(dateOnly, "/")
		day = fmt.Sprintf("%0*s", 2, dt[0])
		month = fmt.Sprintf("%0*s", 2, dt[1])
		year = fmt.Sprintf("%0*s", 2, dt[2])
	}

	//Check date is BR valid
	regex = "^([0-9]{4}-[0-9]{1,2}-[0-9]{1,2})"
	isDb := Matches(date, regex)

	if isDb {
		dt := strings.Split(dateOnly, "-")
		day = fmt.Sprintf("%0*s", 2, dt[2])
		month = fmt.Sprintf("%0*s", 2, dt[1])
		year = fmt.Sprintf("%0*s", 2, dt[0])
	}

	layout := "2006-01-02T15:04:05Z"
	formatted = fmt.Sprintf("%s-%s-%sT00:00:00Z", year, month, day)

	if len(datetime) > 1 {
		dt := strings.Split(datetime[1], ":")
		hour = fmt.Sprintf("%0*s", 2, dt[0])
		min = fmt.Sprintf("%0*s", 2, dt[1])
		sec = fmt.Sprintf("%0*s", 2, dt[2])
		formatted = fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", year, month, day, hour, min, sec)
	}

	dateTime, err := time.Parse(layout, formatted)

	if year == "0001" {
		isValid = false
	} else if err != nil {
		isValid = false
	} else {
		isValid = true
	}

	return
}
