package router

import (
	"hot/utils"
	"net/http"
	"strings"
)

type Router struct {
	get  map[string]routerEntry
	post map[string]routerEntry
}

type routerEntry struct {
	handler http.Handler
	pattern string
}

func NewRouter() *Router { return new(Router) }

var DefaultRouter = &defaultRouter

var defaultRouter Router

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		utils.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h := router.handler(r)
	h.ServeHTTP(w, r)
}

func (router *Router) handler(r *http.Request) http.Handler {
	method := r.Method
	path := r.URL.Path
	switch method {
	case http.MethodGet:
		if routerEntry, ok := router.get[path]; ok {
			return routerEntry.handler
		}
		for _, get := range router.get {
			if strings.HasPrefix(path, get.pattern) {
				return get.handler
			}
		}
		return http.NotFoundHandler()
	case http.MethodPost:
		if routerEntry, ok := router.post[path]; ok {
			return routerEntry.handler
		}
		for _, post := range router.get {
			if strings.HasPrefix(path, post.pattern) {
				return post.handler
			}
		}
		return http.NotFoundHandler()

	default:
		return http.NotFoundHandler()
	}
}

func GetHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultRouter.GetHandleFunc(pattern, handler)
}

func (router *Router) GetHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		utils.Error.Panic("http: nil handler")
	}
	router.GetHandle(pattern, http.HandlerFunc(handler))
}

func (router *Router) GetHandle(pattern string, handler http.Handler) {

	if pattern == "" {
		utils.Error.Panic("http: invalid pattern")
	}
	if handler == nil {
		utils.Error.Panic("http: nil handler")
	}
	if _, exist := router.get[pattern]; exist {
		utils.Error.Panic("http: multiple registrations for " + pattern)
	}

	if router.get == nil {
		router.get = make(map[string]routerEntry)
	}
	r := routerEntry{handler: handler, pattern: pattern}
	router.get[pattern] = r
}

func PostHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	DefaultRouter.PostHandleFunc(pattern, handler)
}

func (router *Router) PostHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		utils.Error.Panic("http: nil handler")
	}
	router.PostHandle(pattern, http.HandlerFunc(handler))
}

func (router *Router) PostHandle(pattern string, handler http.Handler) {

	if pattern == "" {
		utils.Error.Panic("http: invalid pattern")
	}
	if handler == nil {
		utils.Error.Panic("http: nil handler")
	}
	if _, exist := router.post[pattern]; exist {
		utils.Error.Panic("http: multiple registrations for " + pattern)
	}

	if router.post == nil {
		router.post = make(map[string]routerEntry)
	}
	r := routerEntry{handler: handler, pattern: pattern}
	router.post[pattern] = r
}
