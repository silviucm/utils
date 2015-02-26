package utils

import (
	"strings"
)

type SliceUtils struct{}

// single variable acting as the SliceUtils "subpackage" inside the legit utils package
var Slice SliceUtils

// Merges the entire elements of a slice into a string and returns it.
// If the slice is nil, return empty string.
func (dummyReceiver *SliceUtils) ToString(sliceOfStrings []string) (mergedValue string) {

	if sliceOfStrings == nil {
		return ""
	}

	mergedValue = strings.Join(sliceOfStrings, "")

	return mergedValue

}
