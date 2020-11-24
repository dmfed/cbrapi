package cbrcurrency

import (
	"regexp"
	"strings"
	"time"
)

type stringtotime struct {
	time time.Time
	str  string
}

var (
	dateStringFormatForAPIwithSlash = regexp.MustCompile("^\\d{2}/\\d{2}/\\d{4}$")
	dateStringFormatForAPIwithDot   = regexp.MustCompile("^\\d{2}\\.\\d{2}\\.\\d{4}$")
)

// apidate accepts either time.Time type OR date in form of string strictly
// in "DD/MM/YYYY" or "DD.MM.YYY" formats (used by API). If either string is wrongly
// formatted or an unsupported type is passed, apidate will return empty object.
func apidate(input interface{}) *stringtotime {
	var stt stringtotime
	switch input.(type) { // No one likes this, but this seems to be the right place.
	case time.Time:
		stt.time = input.(time.Time)
		stt.str = stt.time.Format("02/01/2006")
	case string:
		if dateStringFormatForAPIwithSlash.MatchString(input.(string)) {
			stt.str = input.(string)
		} else if dateStringFormatForAPIwithDot.MatchString(input.(string)) {
			stt.str = strings.ReplaceAll(input.(string), ".", "/")
		}
		date, err := time.Parse("02/01/2006", stt.str) // Note that this now returns UTC!
		if err == nil {
			stt.time = date
		}
	default:
	}
	return &stt
}

func (stt *stringtotime) timeobject() time.Time {
	return stt.time
}

func (stt *stringtotime) stringobject() string {
	return stt.str
}
