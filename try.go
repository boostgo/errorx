package errorx

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/boostgo/convert"
)

// Try recovers if panic was thrown.
//
// Return error of provided function and recover error
func Try(fn func() error) (err error) {
	defer func() {
		if err == nil {
			err = CatchPanic(recover())
		}
	}()

	return fn()
}

// TryContext is like Try but provided function has context as an argument
func TryContext(ctx context.Context, fn func(ctx context.Context) error) error {
	if ctx == nil {
		ctx = context.Background()
	}

	return Try(func() error {
		return fn(ctx)
	})
}

// TryMust run provided function but ignore error
func TryMust(tryFunc func() error) {
	_ = Try(tryFunc)
}

// CatchPanic got recover() return value and convert it to error
func CatchPanic(err any) error {
	if err == nil {
		return nil
	}

	return New("PANIC RECOVER").
		SetError(errors.New(convert.String(err))).
		AddContext("trace", convert.String(debug.Stack()))
}
