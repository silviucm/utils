package utils

import "strings"

// If the given string is empty, offer an alternative to return in its stead
func StringEmptyAlternative(valueToReturnIfNotEmpty, valueToReturnIfOriginalEmpty string) string {

	if valueToReturnIfNotEmpty == "" {
		return valueToReturnIfOriginalEmpty
	}

	return valueToReturnIfNotEmpty

}

func StringContains(parentString string, stringContainedInParent string) bool {

	return strings.Contains(parentString, stringContainedInParent)

}

func StringReplaceAll(parentString string, oldString string, newString string) string {

	return strings.Replace(parentString, oldString, newString, -1)

}
