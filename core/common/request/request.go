package request

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/lierbai/nspider/core/common/mlog"
)

// Request 请求
type Request struct {
	Method        string //请求方式
	URL           string //处理后的链接
	Header        http.Header
	Cookies       []*http.Cookie
	Postdata      string
	RespType      string
	Urltag        string
	checkRedirect func(req *http.Request, via []*http.Request) error
	Meta          interface{}
}

// NewRequest new
func NewRequest(
	method string, url string, header http.Header, cookies []*http.Cookie, postdata string, respType string, urltag string,
	checkRedirect func(req *http.Request, via []*http.Request) error, meta interface{}) *Request {
	return &Request{method, url, header, cookies, postdata, urltag, respType, checkRedirect, meta}
}

// CopyRequest CopyRequestobject
func (object *Request) CopyRequest() *Request {
	temp := *object
	return &temp
}

// GenHTTPRequest 123
func (object *Request) GenHTTPRequest() (*http.Request, error) {
	// 先生成一个普通的请求.再填充相关的参数
	httpreq, err := http.NewRequest(object.GetMethod(), object.GetURL(), strings.NewReader(object.GetPostdata()))
	if header := object.GetHeader(); header != nil {
		httpreq.Header = object.GetHeader()
	}
	if cookies := object.GetCookies(); cookies != nil {
		for i := range cookies {
			httpreq.AddCookie(cookies[i])
		}
	}
	return httpreq, err
}

// GetURL get
func (object *Request) GetURL() string {
	return object.URL
}

// GetURLTag get
func (object *Request) GetURLTag() string {
	return object.Urltag
}

// GetMethod get
func (object *Request) GetMethod() string {
	return object.Method
}

// GetPostdata get
func (object *Request) GetPostdata() string {
	return object.Postdata
}

// GetHeader get
func (object *Request) GetHeader() http.Header {
	return object.Header
}

// GetCookies get
func (object *Request) GetCookies() []*http.Cookie {
	return object.Cookies
}

// GetResponceType get
func (object *Request) GetResponceType() string {
	return object.RespType
}

// GetRedirectFunc get
func (object *Request) GetRedirectFunc() func(req *http.Request, via []*http.Request) error {
	return object.checkRedirect
}

// GetMeta get
func (object *Request) GetMeta() interface{} {
	return object.Meta
}

// SetURL get
func (object *Request) SetURL(url string) *Request {
	object.URL = url
	return object
}

// SetURLTag get
func (object *Request) SetURLTag(urltag string) *Request {
	object.Urltag = urltag
	return object
}

// SetMethod get
func (object *Request) SetMethod(method string) *Request {
	object.Method = method
	return object
}

// SetPostdata get
func (object *Request) SetPostdata(postdata string) *Request {
	object.Postdata = postdata
	return object
}

// SetHeader get
func (object *Request) SetHeader(header http.Header) *Request {
	object.Header = header
	return object
}

// setHeaderbyFile
func (object *Request) setHeaderbyFile(headerFile string, replace bool) *Request {
	_, err := os.Stat(headerFile)
	if err != nil {
		return object
	}
	b, err := ioutil.ReadFile(headerFile)
	if err != nil {
		mlog.LogInst().LogError(err.Error())
		object.Header = nil
		return object
	}
	json, err := simplejson.NewJson(b)
	object.setHeaderByJSON(json, replace)
	return object
}

// setHeaderByJSON read json file
func (object *Request) setHeaderByJSON(json *simplejson.Json, replace bool) *Request {
	if replace || object.Header == nil {
		object.Header = make(http.Header)
	}
	headerArr, GetErr := json.Get("Header").Map()
	if GetErr != nil {
		mlog.LogInst().LogError(GetErr.Error())
		return object
	}
	for k, v := range headerArr {
		object.Header.Add(k, v.(string))
	}
	return object
}

// SetCookies get
func (object *Request) SetCookies(cookies []*http.Cookie) *Request {
	object.Cookies = cookies
	return object
}

// SetResponceType get
func (object *Request) SetResponceType(respType string) *Request {
	object.RespType = respType
	return object
}

// SetRedirectFunc get
func (object *Request) SetRedirectFunc(checkredirect func(req *http.Request, via []*http.Request) error) *Request {
	object.checkRedirect = checkredirect
	return object
}

// SetMeta get
func (object *Request) SetMeta(meta interface{}) *Request {
	object.Meta = meta
	return object
}
