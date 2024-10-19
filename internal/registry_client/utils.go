package registry_client

import (
	"encoding/base64"
	"errors"
	"net/http"
	"regexp"
)

var (
	// ErrNoMorePages used for cursor pagination state of registry entries
	ErrNoMorePages = errors.New("no more pages")
)

// NOTE: пока не используется, было взято с просторов интернета для пагинации
// getPaginationNextLink extract link for result pagination
// Compliant client implementations should always use the Link header value when proceeding through results linearly.
// The client may construct URLs to skip forward in the catalog.
//
// To get the next result set, a client would issue the request as follows, using the URL encoded in the described Link header:
//
//	GET /v2/_catalog?n=<n from the request>&last=<last repository value from previous response>
//
// The URL for the next block is encoded in RFC 5988 (https://tools.ietf.org/html/rfc5988#section-5)
func getPaginationNextLink(resp *http.Response) (string, error) {
	var nextLinkRE = regexp.MustCompile(`^ *<?([^;>]+)>? *(?:;[^;]*)*; *rel="?next"?(?:;.*)?`)

	for _, link := range resp.Header[http.CanonicalHeaderKey("Link")] {
		parts := nextLinkRE.FindStringSubmatch(link)
		if parts != nil {
			return parts[1], nil
		}
	}
	return "", ErrNoMorePages
}

func name2id(value string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(value))
	return enc
}

func id2name(value string) (string, error) {
	dec, err := base64.StdEncoding.DecodeString(value)
	return string(dec), err
}
