package utils

import "github.com/twinj/uuid"

type GuidUtils struct{}

// single variable acting as the Strings "subpackage" inside the legit utils package
var Guid GuidUtils

// Generates a new UUID and returns it as string
func (dummyReceiver *GuidUtils) New() string {

	return uuid.NewV4().String()

}
