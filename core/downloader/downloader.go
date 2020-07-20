package downloader

import "github.com/lierbai/nspider/core/common/request"

// Downloader interface
type Downloader interface {
	// Download(req *request.Request) *request.Request
	Download(args map[string]string) *request.Request
}





