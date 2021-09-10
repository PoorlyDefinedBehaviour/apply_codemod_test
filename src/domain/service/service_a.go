package service

import "github.com/pkg/errors"

var AError = errors.New("a")

func DoSomethingA() error {
	return errors.Wrapf(AError, "some context")
}
