// Package downloader is the main module of GO_SPIDER for download page.
package downloader

import (
	"github.com/lierbai/nspider/core/common/page"
	"github.com/lierbai/nspider/core/common/request"
)

// Downloader interface
type Downloader interface {
	Downloader(req *request.Request) *page.Page
}
