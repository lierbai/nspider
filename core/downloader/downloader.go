// Package downloader is the main module of GO_SPIDER for download page.
package downloader

// Downloader 下载器
type Downloader interface {
	Download(req string) string
}
