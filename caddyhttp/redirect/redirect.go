// Package redirect is middleware for redirecting certain requests
// to other locations.
package redirect

import (
	"fmt"
	"html"
	"net/http"

	"github.com/orbsmiv/caddy/caddyhttp/httpserver"
)

// Redirect is middleware to respond with HTTP redirects
type Redirect struct {
	Next  httpserver.Handler
	Rules []Rule
}

// ServeHTTP implements the httpserver.Handler interface.
func (rd Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	for _, rule := range rd.Rules {
		if (rule.FromPath == "/" || r.URL.Path == rule.FromPath) && schemeMatches(rule, r) && rule.Match(r) {
			to := httpserver.NewReplacer(r, nil, "").Replace(rule.To)
			if rule.Meta {
				safeTo := html.EscapeString(to)
				fmt.Fprintf(w, metaRedir, safeTo, safeTo)
			} else {
				http.Redirect(w, r, to, rule.Code)
			}
			return 0, nil
		}
	}
	return rd.Next.ServeHTTP(w, r)
}

func schemeMatches(rule Rule, req *http.Request) bool {
	return (rule.FromScheme() == "https" && req.TLS != nil) ||
		(rule.FromScheme() != "https" && req.TLS == nil)
}

// Rule describes an HTTP redirect rule.
type Rule struct {
	FromScheme   func() string
	FromPath, To string
	Code         int
	Meta         bool
	httpserver.RequestMatcher
}

// Script tag comes first since that will better imitate a redirect in the browser's
// history, but the meta tag is a fallback for most non-JS clients.
const metaRedir = `<!DOCTYPE html>
<html>
	<head>
		<script>window.location.replace("%s");</script>
		<meta http-equiv="refresh" content="0; URL='%s'">
	</head>
	<body>Redirecting...</body>
</html>
`
