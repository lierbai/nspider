//
package scheduler_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/lierbai/nspider/core/common/request"
	"github.com/lierbai/nspider/core/scheduler"
)

func TestQueueScheduler(t *testing.T) {
	var r *request.Request
	// 先创建一个标准请求模板
	var h = make(http.Header)
	h.Add("t", "backend.dev.com")
	r = request.NewRequest("GET", "http://backend.dev.com/api/gettest?name=lierbai", h, nil, "{\"action\":\"spider\"}", "json", "", nil, nil)
	var s *scheduler.QueueScheduler
	s = scheduler.NewQueueScheduler(false)
	s.Push(r)
	var count int = s.Count()
	if count != 1 {
		t.Error("count error")
	}
	fmt.Println(count)
	var r1 *request.Request
	r1 = s.Poll()
	if r1 == nil {
		t.Error("poll error")
	}
	fmt.Printf("%v\n", r1)
	client := &http.Client{
		CheckRedirect: r1.GetRedirectFunc(),
	}
	var httpreq, error = r1.GenHTTPRequest()
	if error != nil {
		fmt.Printf("client do error %v \r\n", error)
	}
	var resp *http.Response
	resp, error = client.Do(httpreq)
	if error != nil {
		fmt.Printf("client do error %v \r\n", error)
	}
	bres, err := ioutil.ReadAll(resp.Body)
	res, err := simplejson.NewJson([]byte(bres))
	if err != nil {
		fmt.Printf("json %v\n", err)
	}
	fmt.Printf("%v\n", res)

}
