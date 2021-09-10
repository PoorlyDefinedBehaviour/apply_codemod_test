package service

import (
	"github.com/pkg/errors"
	"fmt"
)

var AError = errors.New("a")

func DoSomethingA() error {
	return errors.Wrapf(AError, "ome contex: %w")
}
