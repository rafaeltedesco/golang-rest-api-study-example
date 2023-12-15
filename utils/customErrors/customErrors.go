package customErrors

import "errors"

var (
	ErrTaskIsDone   = errors.New("Already marked as done")
	ErrNotFound     = errors.New("Not found")
	ErrCannotUndone = errors.New("Cannot undone a not finished task")
)
