package store

type SealedError struct {
	message string
	epoch uint64
	address uint64
}

func (e SealedError) String() string {
	return e.message
}
