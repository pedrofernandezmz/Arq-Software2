package repositories

import (
	"fmt"
	"time"

	"github.com/karlseguin/ccache/v2"
	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
)

type RepositoryCCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewCCache(maxSize int64, itemsToPrune uint32, defaultTTL time.Duration) *RepositoryCCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	fmt.Println("[CCache] Initialized")
	return &RepositoryCCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCCache) Get(id string) (dtos.ItemDTO, e.ApiError) {
	item2 := repo.Client.Get(id)
	if item2 == nil {
		return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	if item2.Expired() {
		return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return item2.Value().(dtos.ItemDTO), nil
}

func (repo *RepositoryCCache) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) Update(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	repo.Client.Set(item.Id, item, repo.DefaultTTL)
	return item, nil
}

func (repo *RepositoryCCache) Delete(id string) e.ApiError {
	repo.Client.Delete(id)
	return nil
}
