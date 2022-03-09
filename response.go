package verificationcode

type Response struct {
	Status Status
	Body   []byte
}

func NewResponse() *Response {
	return &Response{}
}
