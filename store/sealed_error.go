package store

type SealedError struct {
	message string
	epoch uint64
	address uint64
}

func (e SealedError) Error() string {
	return e.message
}
