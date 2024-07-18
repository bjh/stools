package domain

import (
	"net/url"
	"regexp"
	"strings"
)

// StripProtocol will strip an existing protocol from a URL
// only if it is http or https
func StripProtocol(url string) string {
	re := regexp.MustCompile("^https?://")

	return re.ReplaceAllString(url, "")
}

// RemoveProtocol will remove the protocol from a URL
// handles urls that start with // as well
func RemoveProtocol(uri string) string {
	cleaned := strings.Replace(uri, "http://", "", 1)
	cleaned = strings.Replace(cleaned, "https://", "", 1)
	cleaned = strings.Replace(cleaned, "/", "", 20)

	return cleaned
}

// ForceProtocol will force a URL to use a specific protocol
func ForceProtocol(uri string, protocol string) string {
	uri = StripProtocol(uri)
	rx := regexp.MustCompile(`^https?`)
	rx2 := regexp.MustCompile(`^:?//`)

	output := rx.ReplaceAllLiteralString(uri, "")
	output = rx2.ReplaceAllLiteralString(output, "")

	return protocol + "://" + output
}

// HttpsToHttp will force an HTTPS URL to use HTTP
// only if it has an HTTPS protocol
func HttpsToHttp(uri string) string {
	if strings.HasPrefix(uri, "https") {
		return strings.Replace(uri, "https", "http", 1)
	}

	return uri
}

// HttpToHttps will force an HTTP URL to use HTTPS
// only if it has an HTTP protocol
func HttpToHttps(uri string) string {
	if strings.HasPrefix(uri, "http:") {
		return strings.Replace(uri, "http:", "https:", 1)
	}

	return uri
}

// PrependHttps will prepend https to a url if it is not already there
func PrependHttps(uri string) string {
	if strings.HasPrefix(uri, "https:") {
		return uri
	}

	if strings.HasPrefix(uri, "http:") {
		return ForceProtocol(uri, "https")
	}

	uri = RemoveProtocol(uri)

	return "https://" + uri
}

// PrependHttp will prepend http to a url if it is not already there
func PrependHttp(uri string) string {
	if strings.HasPrefix(uri, "http:") {
		return uri
	}

	if strings.HasPrefix(uri, "https:") {
		return ForceProtocol(uri, "http")
	}

	return "http://" + uri
}

// IsValidURI returns whether a URI is valid or not
func IsValidURI(uri string) bool {
	_, err := url.ParseRequestURI(uri)

	return err == nil
}
