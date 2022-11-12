package services

import (
	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) Get(id string) (dtos.ItemDTO, e.ApiError) {
	return dtos.ItemDTO{
		Id:     "12345",
		Titulo: "Mocked item",
	}, nil
}

func (ServiceMock) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	return dtos.ItemDTO{
		Id:     "12345",
		Titulo: item.Titulo,
	}, nil
}
