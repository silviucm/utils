package utils

import "strconv"

type ConvertUtils struct{}

// single variable acting as the ConvertUtils "subpackage" inside the legit utils package
var Convert ConvertUtils

// Converts int64 to string
func (dummyReceiver *ConvertUtils) Int64ToString(integerValue int64) string {

	return strconv.FormatInt(integerValue, 10)

}

// Converts a string to int64
func (dummyReceiver *ConvertUtils) StringToInt64(valueAsString string) (int64, error) {

	return strconv.ParseInt(valueAsString, 10, 64)

}
