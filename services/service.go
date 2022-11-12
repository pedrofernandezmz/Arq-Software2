package services

import (
	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
)

type Service interface {
	Get(id string) (dtos.ItemDTO, e.ApiError)
	Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError)
}
