package HANDLER03

import (
	"GOLANG101/LIBRARY/ip"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

type Handler interface {
	ServeHTTP(c *Context)
}
type HandlerFunc func(ctx *Context)

func (f HandlerFunc) ServeHTTP(ctx *Context) {
	f(ctx)
}
func (engine *Engine) Start() error {
	listener, e := net.Listen("tcp", "127.0.0.1:8080")
	if e != nil {
		log.Print("net.Listen出错", e)
		return e
	}
	log.Print("net.Listen启动成功")
	server := &http.Server{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	//这个函数将会无限期阻止goroutine，直到一个error发生
	//只能用goroutine，否则的话，这个方法就只能运行一个server
	//因为只要第一次调用之后，监听器就会在后面一直死循环了，这个方法就一直没有执行完
	go func() {
		if e := engine.RunServer(server, listener); e != nil {
			log.Print("RunServer出错", e)
			return
		}
	}()
	return nil
}

func (engine *Engine) RunServer(server *http.Server, listener net.Listener) (err error) {
	//engine里面的多路复用器给了server，这个server是在start的时候new的
	server.Handler = engine.mux
	//把这个在start的时候定义的server赋值给engine
	engine.server.Store(server)
	if err = server.Serve(listener); err != nil {
		return
	}
	return
}
func New() *Engine {
	engine := &Engine{
		address:    ip.InternalIP(),
		mux:        http.NewServeMux(),
		metastore:  make(map[string]map[string]interface{}),
		injections: make([]injection, 0),
	}
	engine.addRoute("GET", "/metrics", monitor())
	return engine
}

type Engine struct {
	address string

	mux       *http.ServeMux
	server    atomic.Value
	metastore map[string]map[string]interface{}

	injections []injection
}

type injection struct {
	pattern  *regexp.Regexp
	handlers []HandlerFunc
}

func (engine *Engine) addRoute(method, path string, handlers ...HandlerFunc) {
	if path[0] != '/' {
		panic("blademaster: path must begin with '/'")
	}
	if method == "" {
		panic("blademaster: HTTP method can not be empty")
	}
	if len(handlers) == 0 {
		panic("blademaster: there must be at least one handler")
	}
	if _, ok := engine.metastore[path]; !ok {
		engine.metastore[path] = make(map[string]interface{})
	}
	engine.metastore[path]["method"] = method
	engine.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		c := &Context{
			Context:  nil,
			index:    -1,
			handlers: nil,
			Keys:     nil,
			method:   "",
			Error:    nil,
		}

		c.Request = req
		c.Writer = w
		c.handlers = handlers
		c.method = method

		engine.handleContext(c)
	})
}

func (engine *Engine) handleContext(context *Context) {
	req := context.Request
	ctype := req.Header.Get("Content-Type")
	switch {
	case strings.Contains(ctype, "multipart/form-data"):
		req.ParseMultipartForm(32 << 20)
	default:
		req.ParseForm()
	}
	context.Next()
}
