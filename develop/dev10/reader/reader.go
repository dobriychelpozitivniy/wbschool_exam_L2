package reader

type Reader interface {
	Read([]byte) (int, error)
}
