package util

import (
	"time"
)

var (
	defaultTimeProviderInstance = defaultTimeProvider{}
)

type TimeProvider interface {
	Now() time.Time
}

type defaultTimeProvider struct{}

func GetDefaultTimeProvider() TimeProvider {
	return defaultTimeProviderInstance
}

func (defaultTimeProvider) Now() time.Time {
	return time.Now()
}
