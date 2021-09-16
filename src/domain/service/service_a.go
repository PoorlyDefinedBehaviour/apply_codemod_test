package service

import (
	"fmt"
	"github.com/pkg/errors"
)

var AError = errors.New("a")

func DoSomethingA() error {
	fmt.Errorf("some context: %w", AError)
}
