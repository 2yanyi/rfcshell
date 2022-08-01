package rfcshell

import (
	"errors"
	"github.com/matsuwin/errcause"
	"net/http"
	"strings"
)

func service(handles []handleFunc, w *http.ResponseWriter, r *http.Request) {
	for j := 0; j < len(handles); j++ {
		if err := handles[j](*w, r); err == nil {
			break
		}
	}
}

func (sh ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer errcause.Recover()
	defer r.Body.Close()

	services, ok := multiplexer[r.URL.Path]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte("404"))
		return
	}
	if hook != nil {
		_ = hook(w, r) // requestHookFunction
	}

	// route matching
	paths := strings.Split(r.URL.Path, "/")
	for i := 0; i < len(paths); i++ {
		if i == 0 || i == len(paths)-1 {
			continue
		}
		sh.prefix += "/" + paths[i]
		if handles, has := multiplexer[sh.prefix]; has {
			service(handles, &w, r)
		}
	}
	service(services, &w, r)
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
