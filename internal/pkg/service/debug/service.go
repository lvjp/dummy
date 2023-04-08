package debug

import "os"

var Version = "<dev>"
var BuildTimestamp = "<build-timestamp>"

type Service interface {
	Version() string
	BuildTimestamp() string
	Environment() []string
}

func NewService() Service {
	return serviceImpl{}
}

type serviceImpl struct{}

func (serviceImpl) Version() string {
	return Version
}

func (serviceImpl) BuildTimestamp() string {
	return BuildTimestamp
}

func (serviceImpl) Environment() []string {
	return os.Environ()
}
