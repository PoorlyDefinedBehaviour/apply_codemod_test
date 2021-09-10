package service

import "github.com/pkg/errors"

var BError = errors.New("a")

func DoSomethingB() error {
	return errors.Wrapf(BError, "some context")
}
