package service

import (
	"github.com/pkg/errors"
	"fmt"
)

var BError = errors.New("a")

func DoSomethingB() error {
	return fmt.Printf("some context: %w", BError)
}
