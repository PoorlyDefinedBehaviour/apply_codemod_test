package service

import "github.com/pkg/errors"

var BError = errors.New("a")

func DoSomethingC() error {
	return errors.Wrapf(BError, "some context")
}
