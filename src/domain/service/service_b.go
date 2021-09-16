package service

import (
	"fmt"
	"github.com/pkg/errors"
)

var BError = errors.New("a")

func DoSomethingB() error {
	fmt.Errorf("some context: %w", BError)
}
