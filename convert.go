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

// Converts int32 to string
func (dummyReceiver *convertUtils) Int32ToString(integerValue int32) string {

	return strconv.FormatInt(int64(integerValue), 10)

}

// Converts a string to int32
func (dummyReceiver *convertUtils) StringToInt32(valueAsString string) (int32, error) {

	i64, err := strconv.ParseInt(valueAsString, 10, 32)
	return int32(i64), err

}
