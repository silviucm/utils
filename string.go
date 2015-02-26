package utils

import "strings"

type StringUtils struct{}

// single variable acting as the StringUtils "subpackage" inside the legit utils package
var String StringUtils

// If the given string is empty, offer an alternative to return in its stead
func (dummyReceiver *StringUtils) EmptyAlternative(valueToReturnIfNotEmpty, valueToReturnIfOriginalEmpty string) string {

	if valueToReturnIfNotEmpty == "" {
		return valueToReturnIfOriginalEmpty
	}

	return valueToReturnIfNotEmpty

}

func (dummyReceiver *StringUtils) Contains(parentString string, stringContainedInParent string) bool {

	return strings.Contains(parentString, stringContainedInParent)

}

func (dummyReceiver *StringUtils) ReplaceAll(parentString string, oldString string, newString string) string {

	return strings.Replace(parentString, oldString, newString, -1)

}
