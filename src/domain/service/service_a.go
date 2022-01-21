package service

import "github.com/pkg/errors"

var AError = errors.New("a")

func DoSomethingC() error {
	return errors.Wrapf(AError, "some context")
}
