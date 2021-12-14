package main

import (
	"fmt"
	"github.com/IQ-tech/go-errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return errors.Wrapf(errBadStuff, "some context")
}

func main() {
	fmt.Println(DoSomething())
}
