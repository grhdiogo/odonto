
package ecosystem

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"odonto/internal/infra/config"
)

type ecosystemImpl struct {
	tenant string
}

type ecosystemService interface {
	Store(nick, identifier, secret, contact string, role UserRole) (string, error)
	Access(identifier, secret string) (string, error)
	UpdatePassword(userUid, secret string) error
}

func (r *ecosystemImpl) convertResponseToUID(body []byte) string {
	result := new(storeUserResponse)
	err := json.Unmarshal(body, result)
	if err != nil {
		return ""
	}
	//success
	return result.UID
}

func (r *ecosystemImpl) convertResponseToToken(body []byte) string {
	result := &accessUserResponse{}
	err := json.Unmarshal(body, result)
	if err != nil {
		return ""
	}
	//success
	return result.Token
}

func (r *ecosystemImpl) Store(nick, identifier, secret, contact string, role UserRole) (string, error) {
	//retrieve settings
	settings := config.GetSettings()
	data := &storeUserRequest{
		Nick:       nick,
		GID:        settings.Ecosystem.GID,
		Identifier: identifier,
		Secret:     secret,
		Contact:    contact,
		RoleCode:   string(role),
		Tenant:     r.tenant,
	}
	// marshal
	outbuf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(outbuf)
	err := enc.Encode(data)
	if err != nil {
		return "", err
	}
	// url
	url := fmt.Sprintf("%s/pub/v1/tenant/%s/user", settings.Ecosystem.Host, r.tenant)
	// create request
	req, err := http.NewRequest(http.MethodPost, url, outbuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-version", "0.0.1")
	req.Header.Set("x-app-tenant", r.tenant)
	req.Header.Set("x-app-name", settings.Ecosystem.Name)
	req.Header.Set("x-app-token", settings.Ecosystem.Token)
	req.Header.Set("x-versao", "na")
	// execute send / recv
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//receive the response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//
	if resp.StatusCode != 200 {
		e := convertResponseToError(result)
		if e == nil {
			return "", errors.New("failed to create user")
		}
		return "", &userError{
			Code:    e.Code,
			Message: e.Message,
			MID:     e.MID,
			Version: e.Version,
		}
	}
	// success
	return r.convertResponseToUID(result), err
}

func (r ecosystemImpl) Access(identifier, secret string) (string, error) {
	//retrieve settings
	settings := config.GetSettings()
	//
	data := &accessUserRequest{
		Identifier: identifier,
		Secret:     secret,
		Tenant:     r.tenant,
	}
	// marshal
	outbuf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(outbuf)
	err := enc.Encode(data)
	if err != nil {
		return "", err
	}
	//
	url := fmt.Sprintf("%s/pub/v1/auth", settings.Ecosystem.Host)
	// create request
	req, err := http.NewRequest(http.MethodPost, url, outbuf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-version", "0.0.1")
	req.Header.Set("x-app-tenant", r.tenant)
	req.Header.Set("x-app-name", settings.Ecosystem.Name)
	req.Header.Set("x-app-token", settings.Ecosystem.Token)
	req.Header.Set("x-versao", "na")
	// execute send / recv
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	//receive the response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//
	if resp.StatusCode != 200 {
		e := convertResponseToError(result)
		return "", &userError{
			Code:    e.Code,
			Message: e.Message,
			MID:     e.MID,
			Version: e.Version,
		}
	}
	return r.convertResponseToToken(result), nil
}

func (r ecosystemImpl) UpdatePassword(userUid, secret string) error {
	///retrieve settings
	settings := config.GetSettings()
	data := &updatePasswordRequest{
		Secret:  secret,
		MID:     "12345678910",
		Version: "v1",
	}
	// marshal
	outbuf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(outbuf)
	err := enc.Encode(data)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/priv/v1/tenant/%s/user/%s", settings.Ecosystem.Host, r.tenant, userUid)
	// create request
	req, err := http.NewRequest(http.MethodPut, url, outbuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-version", "0.0.1")
	req.Header.Set("x-app-tenant", r.tenant)
	req.Header.Set("x-app-name", settings.Ecosystem.Name)
	req.Header.Set("x-app-token", settings.Ecosystem.Token)
	req.Header.Set("x-versao", "na")
	// execute send / recv
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//receive the response
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//
	if resp.StatusCode != 200 {
		e := convertResponseToError(result)
		return &userError{
			Code:    e.Code,
			Message: e.Message,
			MID:     e.MID,
			Version: e.Version,
		}
	}
	// success
	return nil
}

func NewEcossytemService(tenant string) ecosystemService {
	return &ecosystemImpl{
		tenant: tenant,
	}
}
