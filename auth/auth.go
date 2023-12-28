package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetApiKey extract an api key from
// the headers of an http request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetApiKey(headers *http.Header) (apiKey string, err error) {

	val := headers.Get("Authorization")

	if val == "" {

		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, "  ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	apiKey = vals[1]
	fmt.Printf("val: %v", apiKey)
	return
}
