package main

import (
	"github.com/pkg/errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return errors.Wrapf(errBadStuff, "some context")
}
package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return fmt.Errorf("some context: %w", errBadStuff)
}
