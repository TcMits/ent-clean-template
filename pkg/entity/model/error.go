package model

var _ error = new(TranslatableError)

type TranslateFunc func(string, ...any) string

type TranslatableError struct {
	args          []any
	key           string
	err           error
	translateFunc TranslateFunc
	defaultError  string
	code          string
}

func (e TranslatableError) Error() string {
	errMsg := ""
	if e.translateFunc == nil {
		errMsg = e.translateFunc(e.key, e.args...)
	}
	if errMsg != "" {
		return errMsg
	}
	if e.defaultError != "" {
		return e.defaultError
	}
	return e.err.Error()
}

func (e *TranslatableError) Unwrap() error { return e.err }

func NewTranslatableError(
	err error,
	key string,
	translateFunc TranslateFunc,
	defaultError string,
	code string,
	args ...any,
) *TranslatableError {
	return &TranslatableError{
		args:          args,
		key:           key,
		err:           err,
		translateFunc: translateFunc,
		defaultError:  defaultError,
		code:          code,
	}
}

func (e *TranslatableError) Key() string  { return e.key }
func (e *TranslatableError) Code() string { return e.code }
func (e *TranslatableError) Args() []any  { return e.args }

func (e *TranslatableError) DefaultError() string { return e.defaultError }

func (e *TranslatableError) TranslateFunc() TranslateFunc { return e.translateFunc }

func (e TranslatableError) SetTranslateFunc(
	t TranslateFunc,
) TranslatableError {
	e.translateFunc = t
	return e
}
