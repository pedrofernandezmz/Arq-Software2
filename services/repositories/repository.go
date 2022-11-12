package repositories

import (
	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	"github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
)

type Repository interface {
	Get(id string) (dtos.ItemDTO, errors.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, errors.ApiError)
}
