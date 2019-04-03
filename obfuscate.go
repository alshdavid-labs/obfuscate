package obfuscate

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

// Encode will take a struct and turn it into
// a url safe string
func Encode(value interface{}) (string, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return EncodeString(string(bytes)), nil
}

// EncodeString takes a string and returns a string that
// is a little harder to read
func EncodeString(value string) string {
	base64encoded := base64.StdEncoding.EncodeToString([]byte(value))
	return url.QueryEscape(base64encoded)
}

// Decode will decode an encoded struct or map and
// unmarshall the values into a struct
func Decode(value string, target interface{}) error {
	decodedValue, err := DecodeString(value)
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(decodedValue), &target)
	return nil
}

// DecodeString takes a previously encoded string and
// makes it readable again
func DecodeString(value string) (string, error) {
	base64Value, err := url.QueryUnescape(value)
	decodedValue, err := base64.StdEncoding.DecodeString(base64Value)
	if err != nil {
		return "", err
	}
	return string(decodedValue), nil
}
