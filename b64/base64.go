// This package defines some base64 encoding/decoding functions.

package b64

import "encoding/base64"

func Encode(message []byte) []byte {
	ret := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(ret, message)
	return ret
}

func Decode(message []byte) ([]byte, error) {
	ret := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	_, err := base64.StdEncoding.Decode(ret, message)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func EncodeString(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func DecodeString(message string) (string, error) {
	ret, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
