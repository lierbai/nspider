//
package page_items

import "github.com/lierbai/nspider/core/common/request"

// PageItems m
type PageItems struct {
	req   *request.Request
	items map[string]string
	skip  bool
}

// NewPageItems returns initialized PageItems object.
func NewPageItems(req *request.Request) *PageItems {
	items := make(map[string]string)
	return &PageItems{req: req, items: items, skip: false}
}

// GetRequest returns request of PageItems
func (this *PageItems) GetRequest() *request.Request {
	return this.req
}

// AddItem saves a KV result into PageItems.
func (this *PageItems) AddItem(key string, item string) {
	this.items[key] = item
}

// GetItem returns value of the key.
func (this *PageItems) GetItem(key string) (string, bool) {
	t, ok := this.items[key]
	return t, ok
}

// GetAll returns all the KVs result.
func (this *PageItems) GetAll() map[string]string {
	return this.items
}

// SetSkip set skip true to make this page not to be processed by Pipeline.
func (this *PageItems) SetSkip(skip bool) *PageItems {
	this.skip = skip
	return this
}

// GetSkip returns skip label.
func (this *PageItems) GetSkip() bool {
	return this.skip
}
