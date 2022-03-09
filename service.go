package verificationcode

type Service interface {
	Response(ctx *Context, refresh bool) (response *Response, err error)
	Challenge(ctx *Context, code []byte) (result *Result, err error)
}
