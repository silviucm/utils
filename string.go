package utils

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stringUtils struct{}

// single variable acting as the StringUtils "subpackage" inside the legit utils package
var String stringUtils

var camelCaseRegularExpression = regexp.MustCompile("[0-9A-Za-z]+")

// If the given string is empty, offer an alternative to return in its stead
func (dummyReceiver *stringUtils) EmptyAlternative(valueToReturnIfNotEmpty, valueToReturnIfOriginalEmpty string) string {

	if valueToReturnIfNotEmpty == "" {
		return valueToReturnIfOriginalEmpty
	}

	return valueToReturnIfNotEmpty

}

func (dummyReceiver *stringUtils) Contains(parentString string, stringContainedInParent string) bool {

	return strings.Contains(parentString, stringContainedInParent)

}

func (dummyReceiver *stringUtils) ReplaceAll(parentString string, oldString string, newString string) string {

	return strings.Replace(parentString, oldString, newString, -1)

}

func (dummyReceiver *stringUtils) CamelCase(original string) string {

	if original == "" {
		return ""
	}

	sections := camelCaseRegularExpression.FindAll([]byte(original), -1)
	for i, v := range sections {
		sections[i] = bytes.Title(v)
	}

	// while returning, make sure to lower the first character
	return dummyReceiver.LowerFirstChar(string(bytes.Join(sections, nil)))

}

func (dummyReceiver *stringUtils) LowerFirstChar(original string) string {

	if original == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(original)
	return string(unicode.ToLower(r)) + original[n:]
}
