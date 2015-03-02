package utils

import "strconv"

type convertUtils struct{}

// single variable acting as the ConvertUtils "subpackage" inside the legit utils package
var Convert convertUtils

// Converts int64 to string
func (dummyReceiver *convertUtils) Int64ToString(integerValue int64) string {

	return strconv.FormatInt(integerValue, 10)

}

// Converts a string to int64
func (dummyReceiver *convertUtils) StringToInt64(valueAsString string) (int64, error) {

	return strconv.ParseInt(valueAsString, 10, 64)

}
