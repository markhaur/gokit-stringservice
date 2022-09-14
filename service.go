package main

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return s, ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
