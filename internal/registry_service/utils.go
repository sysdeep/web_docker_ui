package registry_service

import "encoding/base64"

func name2id(value string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(value))
	return enc
}

func id2name(value string) (string, error) {
	dec, err := base64.StdEncoding.DecodeString(value)
	return string(dec), err
}
