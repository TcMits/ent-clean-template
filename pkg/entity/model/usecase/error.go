package usecase

var _ error = new(UseCaseError)

type UseCaseError struct {
	args         []any
	key          string
	err          error
	defaultError string
	code         string
}

func (e *UseCaseError) Error() string { return e.err.Error() }

func (e *UseCaseError) Unwrap() error { return e.err }

func NewUseCaseError(
	err error, key string, defaultError string, code string, args ...any,
) *UseCaseError {
	return &UseCaseError{
		args:         args,
		key:          key,
		err:          err,
		defaultError: defaultError,
		code:         code,
	}
}

func (e *UseCaseError) Key() string          { return e.key }
func (e *UseCaseError) Code() string         { return e.code }
func (e *UseCaseError) Args() []any          { return e.args }
func (e *UseCaseError) DefaultError() string { return e.defaultError }
