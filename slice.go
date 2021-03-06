package utils

import (
	"strings"
)

type sliceUtils struct{}

// single variable acting as the SliceUtils "subpackage" inside the legit utils package
var Slice sliceUtils

// Merges the entire elements of a slice into a string and returns it.
// If the slice is nil, return empty string.
// If separator is not the empty string, inserts whatever it's passed between the elements when merging
func (dummyReceiver *sliceUtils) ToString(sliceOfStrings []string, separator string) (mergedValue string) {

	if sliceOfStrings == nil {
		return ""
	}

	mergedValue = strings.Join(sliceOfStrings, separator)

	return mergedValue

}
