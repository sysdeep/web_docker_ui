package utils

import "strings"

func ResolveURL(base, local string) string {
	return strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(local, "/")
}
