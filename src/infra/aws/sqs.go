package sqs

import (
	"fmt"
	"github.com/pkg/errors"
)

type EnqueueOptions struct {
}

var ErrEnqueueError = errors.New("oops")

func Enqueue(_ EnqueueOptions) error {
	fmt.Errorf("dont know: %w", ErrEnqueueError)
}
