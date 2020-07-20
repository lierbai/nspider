package request

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lierbai/nspider/core/common/util"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"golang.org/x/net/html/charset"
)

// Request class
type Request struct {
	Method        string //请求方式
	URL           string //处理后的链接
	Payload       string
	RespType      string
	Params        map[string]string
	Header        http.Header
	Cookies       []*http.Cookie
	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// Params        map[string]string

// Response class
type Response struct {
	Status   int
	Content  string
	Header   http.Header
	Cookies  []*http.Cookie
	RespType string
}

// NewRequest new class
func NewRequest(method string, url string, payload string, respType string, params map[string]string,
	header http.Header,
	cookies []*http.Cookie,
	checkRedirect func(req *http.Request, via []*http.Request) error) *Request {
	return &Request{method, url, respType, payload, params, header, cookies, checkRedirect}
}

// params map[string]string,

// QuickRequest qc
func QuickRequest(method string, url string, payload string, respType string, params map[string]string) *Request {
	return NewRequest(method, url, respType, payload, params, nil, nil, nil)
}

// Connect 连接
func (r *Request) Connect(proxy *url.URL, t int) *Response {
	if r.Method == "" {
		r.Method = "GET"
	}
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(r.Payload))
	if r.Params != nil {
		q := req.URL.Query()
		for key, val := range r.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	if r.Header != nil {
		req.Header = r.Header
	}
	if r.Cookies != nil {
		for i := range r.Cookies {
			req.AddCookie(r.Cookies[i])
		}
	}
	transport := &http.Transport{}
	// 操作代理
	if proxy != nil {
		transport.Proxy = http.ProxyURL(proxy)
	}
	client := &http.Client{
		Transport: transport,
	}
	timeout := time.Duration(0) * time.Second
	if t > 0 {
		timeout = time.Duration(t) * time.Second
	}
	client.Timeout = timeout
	resp, err := client.Do(req)
	if err != nil {
		log.Error("request.Conent.Do" + err.Error())
		return nil
	}
	var body string
	var sorbody []byte
	if r.Header.Get("Content-Encoding") == "gzip" {
		body = cCEAutoGS(r.Header.Get("Content-Type"), resp.Body)
	} else {
		if sorbody, err = ioutil.ReadAll(resp.Body); err != nil {
			log.Error("request.Conent.ReadAll" + err.Error())
			body = ""
		} else {
			body = string(sorbody)
		}
	}

	defer resp.Body.Close()
	return &Response{resp.StatusCode, body, resp.Header, resp.Cookies(), r.RespType}
}

// cCEAuto changeCharsetEncodingAuto
func cCEAuto(c string, sor io.ReadCloser) string {
	var err error
	var sorbody []byte
	destReader, err := charset.NewReader(sor, c)
	if err != nil {
		log.Error("request.cCEAuto.NewReader" + err.Error())
		log.Debug("o")
		if sorbody, err = ioutil.ReadAll(sor); err != nil {
			log.Error("request.cCEAuto.ReadAllsor" + err.Error())
			return ""
		}
	} else {
		if sorbody, err = ioutil.ReadAll(destReader); err != nil {
			log.Error("request.cCEAuto.ReadAlldestReader" + err.Error())
			return ""
		}
	}
	bodystr := string(sorbody)
	return bodystr
}

// cCEAutoGS changeCharsetEncodingAutoGzipSupport
func cCEAutoGS(c string, sor io.ReadCloser) string {
	var err error
	gzipReader, err := gzip.NewReader(sor)
	if err != nil {
		log.Error("request.cCEAutoGS.gzipReader" + err.Error())
		return cCEAuto(c, sor)
	}
	defer gzipReader.Close()
	destReader, err := charset.NewReader(gzipReader, c)
	if err != nil {
		log.Error("request.cCEAutoGS.destReader" + err.Error())
		destReader = sor
	}
	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		log.Error("request.cCEAutoGS.ReadAll" + err.Error())
	}
	bodystr := string(sorbody)
	return bodystr
}

func convContent(body string, btype string) interface{} {
	if btype == "json" {
		if gjson.Valid(body) {
			return gjson.Parse(body)
		}
		return util.Load(body)
	}
	return body
}
