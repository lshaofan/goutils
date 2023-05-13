package arrutil

const (
	ErrEmptyArray = "Empty array"
	ErrNotArray   = "Not array"
	ErrorNotFound = "Not found"
)

var (
	FindErrorNotFound = NewArrError(ErrorNotFound)
	FindErrorNotArray = NewArrError(ErrNotArray)
	FindErrorEmpty    = NewArrError(ErrEmptyArray)
)

type ArrError struct {
	Err string
}

func NewArrError(err string) *ArrError {
	return &ArrError{Err: err}
}

func (a ArrError) Error() string {
	return a.Err
}
