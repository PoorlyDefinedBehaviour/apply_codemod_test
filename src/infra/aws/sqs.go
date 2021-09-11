package sqs

import (
	"github.com/pkg/errors"
	"fmt"
)

type EnqueueOptions struct {
}

var ErrEnqueueError = errors.New("oops")

func Enqueue(_ EnqueueOptions) error {
	return fmt.Printf("dont know: %w", ErrEnqueueError)
}
