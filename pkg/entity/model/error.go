package model

import "github.com/nicksnyder/go-i18n/v2/i18n"

var _ error = new(TranslatableError)

type TranslateFunc func(*i18n.Message, ...any) string

type TranslatableError struct {
	args          []any
	err           error
	i18nMessage   *i18n.Message
	translateFunc TranslateFunc
	code          string
}

func (e TranslatableError) Error() string {
	errMsg := ""
	if e.translateFunc != nil && e.i18nMessage != nil {
		errMsg = e.translateFunc(e.i18nMessage, e.args...)
	}
	if errMsg != "" {
		return errMsg
	}
	if e.i18nMessage != nil {
		return e.i18nMessage.Other
	}
	return e.err.Error()
}

func (e *TranslatableError) Unwrap() error { return e.err }

func NewTranslatableError(
	err error,
	i18nMessage *i18n.Message,
	code string,
	translateFunc TranslateFunc,
	args ...any,
) *TranslatableError {
	return &TranslatableError{
		args:          args,
		err:           err,
		i18nMessage:   i18nMessage,
		translateFunc: translateFunc,
		code:          code,
	}
}

func (e *TranslatableError) Code() string { return e.code }

func (e TranslatableError) SetTranslateFunc(
	t TranslateFunc,
) TranslatableError {
	e.translateFunc = t
	return e
}
