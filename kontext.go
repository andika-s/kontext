package kontext

import (
	"context"
	"errors"
)

const (
	errorStore    string = "failed to store context value"
	errorRetreive string = "failed to retreive context value"
	errorParent   string = "failed to reach context parent"
)

// Retreive data from stored context data by key
// and given error when the context parent is invalid
// then assert any data to addressed data
func Retreive[T any](parent context.Context, key any) (any, error) {
	if err := verifyParent(parent); err != nil {
		return nil, err
	}

	value, ok := parent.Value(key).(T)
	if !ok {
		return nil, errors.New(errorRetreive)
	}
	return value, nil
}

// Storing data from current context parent
// with its key and value then checked invalid context parent
func Store(parent context.Context, key, value any) (context.Context, error) {
	if err := verifyParent(parent); err != nil {
		return nil, err
	}

	if key == nil || value == nil {
		return nil, errors.New(errorStore)
	}
	return context.WithValue(parent, key, value), nil
}

func verifyParent(parent context.Context) error {
	if parent == nil {
		return errors.New(errorParent)
	}
	return nil
}
