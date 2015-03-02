package utils

import "github.com/twinj/uuid"

type guidUtils struct{}

// single variable acting as the GuidUtils "subpackage" inside the legit utils package
var Guid guidUtils

// Generates a new UUID and returns it as string
func (dummyReceiver *guidUtils) New() string {

	return uuid.NewV4().String()

}
