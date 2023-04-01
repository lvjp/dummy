package version

var Version = "<dev>"

// ServiceMiddleware is a chainable behavior modifier for VersionService.
type ServiceMiddleware func(VersionService) VersionService

// VersionService provides operations on strings.
type VersionService interface {
	Version() string
}

func NewVersionService() VersionService {
	return versionService{}
}

type versionService struct{}

func (versionService) Version() string {
	return Version
}
