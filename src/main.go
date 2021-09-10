package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return errors.Wrapf(errBadStuff, "ome contex: %w")
}

func main() {
	fmt.Println(DoSomething())
}
