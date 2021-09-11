package service

import (
	"github.com/pkg/errors"
	"fmt"
)

var AError = errors.New("a")

func DoSomethingA() error {
	return fmt.Printf("some context: %w", AError)
}
