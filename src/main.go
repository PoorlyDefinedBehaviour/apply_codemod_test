package main

import (
	"github.com/pkg/errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return errors.Wrapf(errBadStuff, "some context")
}
