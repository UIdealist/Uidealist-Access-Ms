// Since most of external services will make use of the access service, a connector
// package is created to handle the communication between the access service and
// external services easily. The connector package will contain functions to check
// access, grant access, revoke access, and update access.

package connector

import (
	"encoding/json"
	"os"
	"uidealist/app/crud"
)

var CONNECTOR_URL string

func Init() {
	CONNECTOR_URL = os.Getenv("CONNECTOR_URL")
}

func GetConnectorURL() string {
	return CONNECTOR_URL
}

func SetConnectorURL(url string) {
	CONNECTOR_URL = url
}

func CheckAccess(policies *crud.AccessList[interface{}]) (*crud.CheckResponse, error) {
	// Query the connector to check if user is allowed to access resources.

	jsonData, err := json.Marshal(policies)
	if err != nil {
		return nil, err
	}

	res, err := PerformRequest(
		CONNECTOR_URL+"/api/v1/check",
		"POST",
		jsonData,
	)

	if err != nil {
		return nil, err
	}

	checkResponse := new(crud.CheckResponse)
	err = json.Unmarshal(res, checkResponse)
	if err != nil {
		return nil, err
	}

	return checkResponse, nil
}

func GrantAccess(policies *crud.AccessList[string]) (*crud.BaseResponse, error) {
	// Query the connector to grant access to resources.

	jsonData, err := json.Marshal(policies)
	if err != nil {
		return nil, err
	}

	res, err := PerformRequest(
		CONNECTOR_URL+"/api/v1/grant",
		"POST",
		jsonData,
	)

	if err != nil {
		return nil, err
	}

	baseResponse := new(crud.BaseResponse)
	err = json.Unmarshal(res, baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

func RevokeAccess(policies *crud.AccessList[string]) (*crud.BaseResponse, error) {
	// Query the connector to revoke access to resources.

	jsonData, err := json.Marshal(policies)
	if err != nil {
		return nil, err
	}

	res, err := PerformRequest(
		CONNECTOR_URL+"/api/v1/revoke",
		"POST",
		jsonData,
	)

	if err != nil {
		return nil, err
	}

	baseResponse := new(crud.BaseResponse)
	err = json.Unmarshal(res, baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}

func UpdateAccess(policies *crud.AccessListUpdate) (*crud.BaseResponse, error) {
	// Query the connector to update access to resources.

	jsonData, err := json.Marshal(policies)
	if err != nil {
		return nil, err
	}

	res, err := PerformRequest(
		CONNECTOR_URL+"/api/v1/update",
		"POST",
		jsonData,
	)

	if err != nil {
		return nil, err
	}

	baseResponse := new(crud.BaseResponse)
	err = json.Unmarshal(res, baseResponse)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
