package handler

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type RequestLog struct {
	ReqID       string     `json:"req_id"`
	Method      string     `json:"method"`
	URLPath     string     `json:"url_path"`
	BodyRequest url.Values `json:"body_request"`
	Time        time.Time  `json:"time"`
}

func (h *Handler) Logger(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, 10)
	_, _ = rand.Read(data)
	id := fmt.Sprintf("%x", sha1.Sum(data))

	body := r.PostForm
	if body == nil {
		body = r.URL.Query()
	}

	rl := RequestLog{
		ReqID:       id,
		Method:      r.Method,
		URLPath:     r.URL.Path,
		BodyRequest: body,
		Time:        time.Now(),
	}

	b, _ := json.Marshal(&rl)

	fmt.Println(string(b))
}
