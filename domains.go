package bandwidth

import (
	"net/http"
	"fmt"
)

const domainsPath = "domains"

// GetDomains returns  a list of the domains that have been created
func (api *Client) GetDomains() ([]map[string] interface{}, error){
	result, _, err :=  api.makeRequest(http.MethodGet, api.concatUserPath(domainsPath), nil, []map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return result.([]map[string] interface{}), nil
}

// CreateDomain creates a new domain
func (api *Client) CreateDomain(data map[string]interface{}) (string, error){
	_, headers, err :=  api.makeRequest(http.MethodPost, api.concatUserPath(domainsPath), data)
	if err != nil {
		return "", err
	}
	return getIDFromLocationHeader(headers), nil
}

// DeleteDomain removes a domain
func (api *Client) DeleteDomain(id string) error{
	_, _, err :=  api.makeRequest(http.MethodDelete, fmt.Sprintf("%s/%s", api.concatUserPath(domainsPath), id))
	return err
}