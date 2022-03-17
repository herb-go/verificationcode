package verificationcode

type Service interface {
	Type() string
	Challenge(ctx *Context) (challenge *Challenge, err error)
	Response(ctx *Context, code []byte) (result *Result, err error)
}
