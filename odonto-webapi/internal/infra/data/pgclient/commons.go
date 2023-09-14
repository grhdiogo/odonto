package pgclient

import (
	"encoding/json"
	"fmt"
)

//===================================================================
//consts
//===================================================================

//to access
const (
	urlPathAccess = "pub/v1/auth"
)

//===================================================================
//structs to access
//===================================================================

type accessUserRequest struct {
	Identifier string `yaml:"identifier"`
	Secret     string `json:"secret"`
	Tenant     string `json:"tenant"`
}

type accessUserResponse struct {
	Token string `json:"jwtToken"`
	MID   string `json:"_mid"`
}

//===================================================================
//structs to users
//===================================================================

//user error
type userError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	MID     string `json:"_mid"`
	Version string `json:"_v"`
}

// store user req
type storeUserRequest struct {
	Nick       string `json:"nick"`
	GID        string `json:"gid"`
	Identifier string `json:"identifier"`
	Secret     string `json:"secret"`
	Contact    string `json:"contact"`
	Tenant     string `json:"tenant"`
	RoleCode   string `json:"roleCode"`
}

//store user resp
type storeUserResponse struct {
	UID string `json:"uid"`
	MID string `json:"_mid"`
}

//remove user req
type removeUserRequest struct {
	UID     string `json:"uid"`
	Tenant  string `json:"tenant"`
	MID     string `json:"_mid"`
	Version string `json:"_v"`
}

// =====================================================================
// struct to update pswd
// =====================================================================

// update pswd
type updatePasswordRequest struct {
	Secret  string `json:"secret"`
	MID     string `json:"_mid"`
	Version string `json:"_v"`
}

// =====================================================================
// Error convert error response
// =====================================================================

func convertResponseToError(body []byte) *userError {
	//
	result := &userError{}
	err := json.Unmarshal(body, result)
	if err != nil {
		return nil
	}
	return result
}

func (e *userError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, mid: %s, version: %s",
		e.Code, e.Message, e.MID, e.Version,
	)
}
