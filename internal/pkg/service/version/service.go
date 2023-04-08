package version

var Version = "<dev>"

type Service interface {
	Version() string
}

func NewService() Service {
	return serviceImpl{}
}

type serviceImpl struct{}

func (serviceImpl) Version() string {
	return Version
}
