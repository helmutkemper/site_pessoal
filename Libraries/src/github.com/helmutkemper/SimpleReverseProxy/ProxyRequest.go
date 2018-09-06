package SimpleReverseProxy

import "net/http"

type ProxyRequest struct {
	R             *http.Request
	QueryString   map[string][]string
	ExpRegMatches map[string]string
	SubDomain     string
	Domain        string
	Port          string
	Path          string
}
