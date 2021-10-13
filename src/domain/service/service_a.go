package service

import (
	"fmt"
	"github.com/pkg/errors"
)

var AError = errors.New("a")

func DoSomethingA() error {
	return fmt.Errorf("some context: %w", AError)
}
