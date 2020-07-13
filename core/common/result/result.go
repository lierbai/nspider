package result

import "net/http"

// Result 结果集
type Result struct {
	status   int
	header   http.Header
	cookies  []*http.Cookie
	response string
}

// NewResult new
func NewResult(s int, h http.Header, c []*http.Cookie, r string) *Result {
	return &Result{s, h, c, r}
}

// GetResponse getbody
func (object *Result) GetResponse() string {
	return object.response
}
