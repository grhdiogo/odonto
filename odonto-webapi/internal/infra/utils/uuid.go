package utils

import uuid "github.com/satori/go.uuid"

//GenerateUUID generate access token
func GenerateUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}
