package service

import (
	"fmt"
	"github.com/pkg/errors"
)

var BError = errors.New("a")

func DoSomethingB() error {
	return fmt.Errorf("some context: %w", BError)
}
