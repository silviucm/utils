package utils

import "strconv"

// Converts int64 to string
func ConvertInt64ToString(integerValue int64) string {

	return strconv.FormatInt(integerValue, 10)

}

// Converts a string to int64
func ConvertStringToInt64(valueAsString string) (int64, error) {

	return strconv.ParseInt(valueAsString, 10, 64)

}
