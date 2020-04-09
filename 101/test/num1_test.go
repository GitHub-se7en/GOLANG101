package test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("给出测试环境需要的数据")
	{
		t.Logf("\t当检查 \"%s\" 状态码是\"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\t应该能够做出Get请求啊", ballotX, err)
			}
			t.Log("\t\t应该能够执行Get请求", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t应该接收到状态码\"%d\".%v", statusCode, checkMark)
			} else {
				t.Errorf("\t\t应该接受\"%d\".%v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
