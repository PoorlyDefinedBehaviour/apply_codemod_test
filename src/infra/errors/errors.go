package errors

import stderrors "errors"

func New(msg string) error {
	return stderrors.New(msg)
}

// pretend we are wrapping the error
func Wrap(err error) error {
	return err
}

// pretend we are wrapping the error
func Wrapf(err error, _ string, _ ...[]interface{}) error {
	return err
}
