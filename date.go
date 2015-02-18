package utils

import "time"

const (
	// See http://golang.org/pkg/time/#Parse
	comparisonTimeFormat = "2006-01-02 15:04:05 MST"
)

// Returns true if secondDateTimeString is after the firstDateTimeString.
// Both dates must be in this format: "YYYY-MM-DD HH:MM:SS" (e.g. 2014-12-22 18:24:43)
func DateIsSecondAfterFirst(firstDateTimeString, secondDateTimeString string) (bool, error) {

	if parsedFirstDate, err := time.Parse(comparisonTimeFormat, firstDateTimeString); err == nil {

		// parse the second date
		if parsedSecondDate, errSecond := time.Parse(comparisonTimeFormat, secondDateTimeString); errSecond == nil {

			// compare the dates
			delta := parsedSecondDate.Sub(parsedFirstDate)

			if delta > 0 {
				return true, nil
			}

			return false, nil

		} else {
			return false, errSecond
		}

	} else {
		return false, err
	}

}

// Supplied date must be in this format: "YYYY-MM-DD HH:MM:SS" (e.g. 2014-12-22 18:24:43)
func DateIsNowAfter(dateTimeString string) (bool, error) {

	if dateTimeParsed, err := time.Parse(comparisonTimeFormat, dateTimeString); err == nil {

		// compare the dates
		delta := time.Now().Sub(dateTimeParsed)

		if delta > 0 {
			return true, nil
		}
		return false, nil

	} else {
		return false, err
	}
}
