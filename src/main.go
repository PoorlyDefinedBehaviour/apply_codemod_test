package main

import (
	"apply_codemod_test/src/infra/errors"
	"fmt"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return errors.Wrapf(errBadStuff, "some context")
}

func main() {
	fmt.Println(DoSomething())
}
