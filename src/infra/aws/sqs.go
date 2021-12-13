package sqs

import "apply_codemod_test/src/infra/errors"

type EnqueueOptions struct {
}

var ErrEnqueueError = errors.New("oops")

func Enqueue(_ EnqueueOptions) error {
	return errors.Wrapf(ErrEnqueueError, "dont know")
}
