package verificationcode

type ResponseType int

const ResponseTypeText = ResponseType(0)
const ResponseTypeBinary = ResponseType(1)
const ResponseTypeJSON = ResponseType(2)
const ResponseTypeJPEG = ResponseType(3)
const ResponseTypePNG = ResponseType(4)

type Response struct {
	Status Status
	Type   ResponseType
	Body   []byte
}

func (r *Response) IsSuccess() bool {
	return r.Status == StatusSuccess
}

func NewResponse() *Response {
	return &Response{}
}
