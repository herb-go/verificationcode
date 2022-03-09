package httpverificationcode

import (
	"net/http"
	"strings"

	"github.com/herb-go/verificationcode"
)

type Engine struct {
	Field      string
	StatusFail int
	Session    Session
	Service    verificationcode.Service
}

func (e *Engine) CreateStore(req *http.Request) *Store {
	return &Store{
		req:     req,
		session: e.Session,
		field:   e.Field,
	}
}
func (e *Engine) CreateContext(r *http.Request, user string) *verificationcode.Context {
	var m = make(map[string]string, len(r.Header))
	for k := range r.Header {
		m[strings.ToLower(k)] = r.Header.Get(k)
	}
	return verificationcode.CreateContext(
		user,
		e.CreateStore(r),
		verificationcode.ClientTypeWeb.CreateClient(r.RemoteAddr, r.Context()).MergeMeta(m),
	)
}

func (e *Engine) ActionResponse(w http.ResponseWriter, r *http.Request) {
	ctx := e.CreateContext(r, "")
	resp, err := e.Service.Response(ctx, false)
	if err != nil {
		panic(err)
	}
	if !resp.IsSuccess() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	contenttype := ""
	switch resp.Type {
	case verificationcode.ResponseTypeText:
		contenttype = "text/plain"
	case verificationcode.ResponseTypeBinary:
		contenttype = "application/octet-stream"
	case verificationcode.ResponseTypeJSON:
		contenttype = "application/json"
	case verificationcode.ResponseTypeJPEG:
		contenttype = "image/jpeg"
	case verificationcode.ResponseTypePNG:
		contenttype = "image/png"
	}
	if contenttype != "" {
		w.Header().Set("Content-Type", contenttype)
	}
	_, err = w.Write(resp.Body)
	if err != nil {
		panic(err)
	}
}
