package caddyhttp

import (
	// plug in the server
	_ "github.com/orbsmiv/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/orbsmiv/caddy/caddyhttp/basicauth"
	_ "github.com/orbsmiv/caddy/caddyhttp/bind"
	_ "github.com/orbsmiv/caddy/caddyhttp/browse"
	_ "github.com/orbsmiv/caddy/caddyhttp/errors"
	_ "github.com/orbsmiv/caddy/caddyhttp/expvar"
	_ "github.com/orbsmiv/caddy/caddyhttp/extensions"
	_ "github.com/orbsmiv/caddy/caddyhttp/fastcgi"
	_ "github.com/orbsmiv/caddy/caddyhttp/gzip"
	_ "github.com/orbsmiv/caddy/caddyhttp/header"
	_ "github.com/orbsmiv/caddy/caddyhttp/index"
	_ "github.com/orbsmiv/caddy/caddyhttp/internalsrv"
	_ "github.com/orbsmiv/caddy/caddyhttp/limits"
	_ "github.com/orbsmiv/caddy/caddyhttp/log"
	_ "github.com/orbsmiv/caddy/caddyhttp/markdown"
	_ "github.com/orbsmiv/caddy/caddyhttp/mime"
	_ "github.com/orbsmiv/caddy/caddyhttp/pprof"
	_ "github.com/orbsmiv/caddy/caddyhttp/proxy"
	_ "github.com/orbsmiv/caddy/caddyhttp/push"
	_ "github.com/orbsmiv/caddy/caddyhttp/redirect"
	_ "github.com/orbsmiv/caddy/caddyhttp/requestid"
	_ "github.com/orbsmiv/caddy/caddyhttp/rewrite"
	_ "github.com/orbsmiv/caddy/caddyhttp/root"
	_ "github.com/orbsmiv/caddy/caddyhttp/status"
	_ "github.com/orbsmiv/caddy/caddyhttp/templates"
	_ "github.com/orbsmiv/caddy/caddyhttp/timeouts"
	_ "github.com/orbsmiv/caddy/caddyhttp/websocket"
	_ "github.com/orbsmiv/caddy/startupshutdown"
)
