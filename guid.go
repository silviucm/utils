package utils

import "github.com/twinj/uuid"

// Generates a new UUID and returns it as string
func GuidGenerateNew() string {

	return uuid.NewV4().String()

}
