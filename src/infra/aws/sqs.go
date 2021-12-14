package sqs

import "github.com/IQ-tech/go-errors"

type EnqueueOptions struct {
}

var ErrEnqueueError = errors.New("oops")

func Enqueue(_ EnqueueOptions) error {
	return errors.Wrapf(ErrEnqueueError, "dont know")
}
