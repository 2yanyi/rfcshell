package rfcshell

import (
	"errors"
	"fmt"
	"github.com/utilgo/errcause"
	"net/http"
	"strings"
)

func (sh serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer errcause.Recover()

	// route check
	if _, has := multiplexer[r.URL.Path]; !has {
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}

	// route front hook
	if sh.FrontHook != nil {
		_ = sh.FrontHook(w, r)
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

func (sh *serverHandler) Route(url string, handles ...handleFunc) *serverHandler {
	prefix := sh.prefix + url
	multiplexer[prefix] = handles
	return &serverHandler{prefix: prefix}
}

func New(addr string) *serverHandler {
	fmt.Printf("bind %s\n", addr)
	return &serverHandler{Server: &http.Server{Addr: addr, Handler: new(serverHandler)}}
}

type handleFunc func(w http.ResponseWriter, r *http.Request) error

type serverHandler struct {
	Server    *http.Server
	prefix    string
	FrontHook handleFunc
}

var multiplexer = make(map[string][]handleFunc)

var Next = errors.New("")
