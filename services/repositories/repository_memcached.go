package repositories

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
	json "github.com/json-iterator/go"
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

	if err := repo.Client.Set(&memcache.Item{ //?????
		Key:   item.Id,
		Value: bytes,
	}); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}

	return item, nil
}

func (repo *RepositoryMemcached) Update(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	bytes, err := json.Marshal(item)
	if err != nil {
		return dtos.ItemDTO{}, e.NewBadRequestApiError(fmt.Sprintf("invalid item %s: %v", item.Id, err))
	}

	if err := repo.Client.Set(&memcache.Item{ //?????
		Key:   item.Id,
		Value: bytes,
	}); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item %s", item.Id), err)
	}

	return item, nil
}

func (repo *RepositoryMemcached) Delete(id string) e.ApiError {
	err := repo.Client.Delete(id)
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}