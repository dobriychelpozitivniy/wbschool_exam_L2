package router

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type hand func(w http.ResponseWriter, r *http.Request)

type Router struct {
	mu          *sync.Mutex
	methods     map[string]map[string]hand
	middlewares []hand
}

func NewRouter() *Router {
	return &Router{mu: &sync.Mutex{}, methods: map[string]map[string]hand{}, middlewares: make([]hand, 0)}
}

func (rr Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	m := r.Method

	switch r.Method {
	case "GET":
		rr.handle(m, url, w, r)
	case "POST":
		rr.handle(m, url, w, r)
	default:
		rr.RespondWithError(w, http.StatusMethodNotAllowed, "invalid method")
	}
}

func (rr *Router) RunServer(port string) {
	http.Handle("/", rr)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (rr *Router) handle(method string, url string, w http.ResponseWriter, r *http.Request) {
	hand := rr.methods[method][url]
	if hand == nil {
		rr.RespondWithError(w, http.StatusNotFound, "not found")

		return
	}

	hand(w, r)

	for _, f := range rr.middlewares {
		f(w, r)
	}
}

func (rr *Router) RespondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (rr *Router) RespondWithError(w http.ResponseWriter, code int, msg string) {
	resp := map[string]string{"error": msg}
	rr.RespondWithJSON(w, code, resp)
}

func (rr *Router) RegisterMethod(method string, url string, handler hand) {
	rr.mu.Lock()
	if rr.methods[method] == nil {
		rr.methods[method] = make(map[string]hand)
	}
	rr.methods[method][url] = handler
	rr.mu.Unlock()
}

func (rr *Router) RegisterMiddleware(handler hand) {
	rr.mu.Lock()
	rr.middlewares = append(rr.middlewares, handler)
	rr.mu.Unlock()
}
