package repositories

import (
	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	"github.com/pedrofernandezmz/Arq-Software2/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.ItemDTO, errors.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, errors.ApiError)
}
