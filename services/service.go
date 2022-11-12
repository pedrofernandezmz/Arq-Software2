package services

import (
	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
)

type Service interface {
	Get(id string) (dtos.ItemDTO, e.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError)
}
