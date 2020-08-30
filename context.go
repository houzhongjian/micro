package micro

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Host           string
	IP             string
}
