package string

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService.
//type Middleware func(Service) Service

// StringService provides operations on strings.
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
}

func NewService() Service {
	return serviceImpl{}
}

type serviceImpl struct{}

func (serviceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (serviceImpl) Count(s string) int {
	return len(s)
}
