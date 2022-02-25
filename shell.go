package rfcshell

import (
	"errors"
	"net/http"
	"strings"

	"github.com/utilgo/errcause"
)

func (sh ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer errcause.Recover()
	defer r.Body.Close()

	if _, has := multiplexer[r.URL.Path]; !has {
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}

	if hook != nil {
		_ = hook(w, r) // requestHookFunction
	}

	// route matching
	for i, prefix := range strings.Split(r.URL.Path, "/") {
		if i != 0 {
			sh.prefix += "/" + prefix
			if handles, has := multiplexer[sh.prefix]; has {
				for _, handle := range handles {
					if err := handle(w, r); err == nil {
						break
					}
				}
			}
		}
	}
}

func (sh *ServerHandler) Route(url string, handles ...handleFunc) *ServerHandler {
	prefix := sh.prefix + url
	multiplexer[prefix] = handles
	return &ServerHandler{prefix: prefix}
}

func New(addr string) *ServerHandler {
	return &ServerHandler{Server: &http.Server{Addr: addr, Handler: new(ServerHandler)}}
}

func SetRequestHookFunction(handle handleFunc) {
	hook = handle
}

type ServerHandler struct {
	Server *http.Server
	prefix string
}

type handleFunc func(w http.ResponseWriter, r *http.Request) error

var multiplexer = make(map[string][]handleFunc)

var hook handleFunc

var Next = errors.New("")
