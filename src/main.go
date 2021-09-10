package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var errBadStuff = errors.New("something happened")

func DoSomething() error {
	return fmt.Errorf("some context: %w",

		"some context: %w")

}

func main() {
	fmt.Println(DoSomething())
}
