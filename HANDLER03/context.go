package HANDLER03

import (
	"context"
	"net/http"
)

type Context struct {
	//既然已经有了这个上下文，为什么还要封装一层
	context.Context

	//为什么这个一定是使用request的指针呢？
	Request *http.Request
	//这个是个接口
	Writer http.ResponseWriter

	index int8
	//这个HandlerFunc是一个函数，函数
	handlers []HandlerFunc

	//根据英文的翻译，这个map应该是针对每一个request都有一个
	Keys map[string]interface{}

	Error error

	method string
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
// See example in godoc.
func (c *Context) Next() {
	c.index++
	s := int8(len(c.handlers))
	for ; c.index < s; c.index++ {
		// only check method on last handler, otherwise middlewares
		// will never be effected if request method is not matched
		if c.index == s-1 && c.method != c.Request.Method {
			code := http.StatusMethodNotAllowed
			http.Error(c.Writer, http.StatusText(code), code)
			return
		}

		c.handlers[c.index](c)
	}
}
