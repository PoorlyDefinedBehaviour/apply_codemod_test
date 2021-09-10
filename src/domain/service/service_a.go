package service

import (
	"github.com/pkg/errors"
	"fmt"
)

var AError = errors.New("a")

func DoSomethingA() error {
	return fmt.Errorf("some context: %w",

		"some context: %w",
	)

}
