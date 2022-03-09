package verificationcode

type Store interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
	Clean() error
}
