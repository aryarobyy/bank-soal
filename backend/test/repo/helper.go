package repo

import "github.com/stretchr/testify/mock"

func GetReturn[T any](args mock.Arguments) (T, error) {
	var zero T

	if args.Get(0) == nil {
		return zero, args.Error(1)
	}

	return args.Get(0).(T), args.Error(1)
}
