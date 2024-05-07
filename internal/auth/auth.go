package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey() extracts the API key from the headers of an HTTP request.
// Example:
// Authorization: APIKey {insert apikey here}

func GetApiKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid auth header")
	}
	if vals[0] != "APIKey" {
		return "", errors.New("invalid auth type")
	}
	return vals[1], nil

}
