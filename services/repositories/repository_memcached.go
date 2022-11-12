package repositories

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	json "github.com/json-iterator/go"
	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
)

type RepositoryMemcached struct {
	Client *memcache.Client
}

func NewMemcached(host string, port int) *RepositoryMemcached {
	client := memcache.New(fmt.Sprintf("%s:%d", host, port))
	fmt.Println("[Memcached] Initialized connection")
	return &RepositoryMemcached{
		Client: client,
	}
}

func (repo *RepositoryMemcached) Get(id string) (dtos.ItemDTO, e.ApiError) {
	item2, err := repo.Client.Get(id)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
		}
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	var itemDTO dtos.ItemDTO
	if err := json.Unmarshal(item2.Value, &itemDTO); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	return itemDTO, nil
}

func (repo *RepositoryMemcached) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	bytes, err := json.Marshal(item)
	if err != nil {
		return dtos.ItemDTO{}, e.NewBadRequestApiError(err.Error())
	}

	if err := repo.Client.Set(&memcache.Item{ //a chequear, buscar donde esta declarado para cambiar
		Key:   item.Id,
		Value: bytes,
	}); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}

	return item, nil
}
