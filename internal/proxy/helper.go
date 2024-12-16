package proxy

import (
	"net/http"
	"strings"
)

// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
// these headers will not get forwarded at all
var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"proxy-Authenticate",
	"proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

// copy headers excluding hop-by-hop headers
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		if !isHopHeader(k) {
			for _, v := range vv {
				dst.Add(k, v)
			}
		}
	}
}

func isHopHeader(header string) bool {
	for _, h := range hopHeaders {
		if h == header {
			return true
		}
	}
	return false
}

// appendHostToXForwardHeader appends host to X-Forwarded-For header
func appendHostToXForwardHeader(header http.Header, host string) {
	// If we aren't the first proxy retain prior
	// X-Forwarded-For information as a comma+space
	// separated list and fold multiple headers into one.
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}
	header.Set("X-Forwarded-For", host)
}
