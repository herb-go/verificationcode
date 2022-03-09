package verificationcode

type Store interface {
	Save(data []byte) error
	Load() ([]byte, error)
	Clean() error
}
