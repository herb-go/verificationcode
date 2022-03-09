package verificationcode

import (
	"fmt"
)

type Result struct {
	Status Status
	Body   []byte
}

func (r *Result) Error() string {
	return fmt.Sprintf("verification code fail:[%d] %s", r.Status, string(r.Body))
}
