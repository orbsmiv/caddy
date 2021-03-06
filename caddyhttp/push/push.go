package push

import (
	"net/http"

	"github.com/orbsmiv/caddy/caddyhttp/httpserver"
)

type (
	// Rule describes conditions on which resources will be pushed
	Rule struct {
		Path      string
		Resources []Resource
	}

	// Resource describes resource to be pushed
	Resource struct {
		Path   string
		Method string
		Header http.Header
	}

	// Middleware supports pushing resources to clients
	Middleware struct {
		Next  httpserver.Handler
		Rules []Rule
	}

	ruleOp func([]Resource)
)
