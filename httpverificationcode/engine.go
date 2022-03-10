package httpverificationcode

import (
	"net/http"
	"strings"

	"github.com/herb-go/verificationcode"
)

const DefaultStatusFail = 422

type Engine struct {
	Field      string
	StatusFail int
	Session    Session
	Service    verificationcode.Service
}

func (e *Engine) NewOptions() *verificationcode.Options {
	return verificationcode.NewOptions()
}
func (e *Engine) CreateStore(req *http.Request) *Store {
	return &Store{
		req:     req,
		session: e.Session,
		field:   e.Field,
	}
}
func (e *Engine) CreateContext(r *http.Request, opt *verificationcode.Options) *verificationcode.Context {
	var m = make(map[string]string, len(r.Header))
	for k := range r.Header {
		m[strings.ToLower(k)] = r.Header.Get(k)
	}
	return verificationcode.CreateContext(
		e.CreateStore(r),
		verificationcode.ClientTypeWeb.CreateClient(r.RemoteAddr, r.Context()).MergeMeta(m),
		opt,
	)
}

func (e *Engine) ActionChallenge(w http.ResponseWriter, r *http.Request) {
	opt := e.NewOptions()
	ctx := e.CreateContext(r, opt)
	resp, err := e.Service.Challenge(ctx)
	if err != nil {
		panic(err)
	}

	contenttype := ""
	switch resp.Type {
	case verificationcode.ChallengeTypeText:
		contenttype = "text/plain"
	case verificationcode.ChallengeTypeBinary:
		contenttype = "application/octet-stream"
	case verificationcode.ChallengeTypeJSON:
		contenttype = "application/json"
	case verificationcode.ChallengeTypeJPEG:
		contenttype = "image/jpeg"
	case verificationcode.ChallengeTypePNG:
		contenttype = "image/png"
	}
	if contenttype != "" {
		w.Header().Set("Content-Type", contenttype)
	}
	if !resp.IsSuccess() {
		status := e.StatusFail
		if status == 0 {
			status = DefaultStatusFail
		}
		w.WriteHeader(status)
	}
	_, err = w.Write(resp.Body)
	if err != nil {
		panic(err)
	}
}

func (e *Engine) ResponseRequest(req *http.Request, opt *verificationcode.Options, code []byte) (result *verificationcode.Result, err error) {
	ctx := e.CreateContext(req, opt)
	return e.Service.Response(ctx, code)
}
