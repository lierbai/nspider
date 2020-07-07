package downloader

import (
	"net/http"

	"github.com/lierbai/nspider/core/common/mlog"
	"github.com/lierbai/nspider/core/common/page"
	"github.com/lierbai/nspider/core/common/request"
)

// HttpDownloader s
type HttpDownloader struct {
}

func NewHttpDownloader() *HttpDownloader {
	return &HttpDownloader{}
}

// Download Download
func (object *HttpDownloader) Download(req *request.Request) *page.Page {

	var mtype string
	p := page.NewPage(req)
	mtype = req.GetResponceType()
	switch mtype {
	case "json":
		fallthrough
	case "jsonp":
		return object.downloadJSON(p, req)
	}
	// 	case "html":
	// 		return this.downloadHtml(p, req)

	// 	case "text":
	// 		return this.downloadText(p, req)
	// 	default:
	// 		mlog.LogInst().LogError("error request type:" + mtype)
	// 	}
	return p
}
func connectByHttp(p *page.Page, req *request.Request) (*http.Response, error) {

}

// downloadFile downloadFile
func (object *HttpDownloader) downloadFile(p *page.Page, req *request.Request) (*page.Page, string) {
	var err error
	var urlstr string
	if urlstr = req.GetURL(); len(urlstr) == 0 {
		mlog.LogInst().LogError("url is empty")
		p.SetStatus(true, "url is empty")
		return p, ""
	}
	var resp *http.Response

	if err != nil {
		return p, ""
	}
	p.SetHeader(resp.Header)
	p.SetCookies(resp.Cookies())

	// get converter to utf-8
	var bodyStr string
	// if resp.Header.Get("Content-Encoding") == "gzip" {
	// 	bodyStr = this.changeCharsetEncodingAutoGzipSupport(resp.Header.Get("Content-Type"), resp.Body)
	// } else {
	// 	bodyStr = this.changeCharsetEncodingAuto(resp.Header.Get("Content-Type"), resp.Body)
	// }
	//fmt.Printf("utf-8 body %v \r\n", bodyStr)
	defer resp.Body.Close()
	return p, bodyStr
}

// downloadJSON download and JSON
func (object *HttpDownloader) downloadJSON(p *page.Page, req *request.Request) *page.Page {
	var err error
	p, destbody := this.downloadFile(p, req)
	if !p.IsSucc() {
		return p
	}
	return p
}
