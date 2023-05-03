package fortune

import "fmt"

type Service interface {
	Create(fortune string) (uuid string, err error)
	Read(uuid string) (fortune string, err error)
	Update(uuid, fortune string) (err error)
	Delete(uuid string) (err error)
}

func NewService() Service {
	return serviceImpl{}
}

type serviceImpl struct{}

func (serviceImpl) Create(fortune string) (string, error) {
	return "", fmt.Errorf("not yet implemented")
}

func (serviceImpl) Read(uuid string) (string, error) {
	return "", fmt.Errorf("not yet implemented")
}

func (serviceImpl) Update(uuid, fortune string) error {
	return fmt.Errorf("not yet implemented")
}

func (serviceImpl) Delete(uuid string) error {
	return fmt.Errorf("not yet implemented")
}
