package httpverificationcode

import "net/http"

var DefaultSessionField = "verificationcode"

type Session interface {
	// Set set session by field name with given value.
	Set(r *http.Request, fieldname string, data []byte) error
	//Get get session by field name with given value.
	Get(r *http.Request, fieldname string) ([]byte, error)
	// Del del session value by field name .
	Del(r *http.Request, fieldname string) error
	// IsNotFoundError check if given error is session not found error.
	IsNotFoundError(err error) bool
}

type Store struct {
	session Session
	req     *http.Request
	field   string
}

func (s *Store) Save(data []byte) error {
	return s.session.Set(s.req, s.field, data)
}
func (s *Store) Load() ([]byte, error) {
	return s.session.Get(s.req, s.field)
}
func (s *Store) Clean() error {
	return s.session.Del(s.req, s.field)
}
