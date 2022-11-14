package services

import (
	"fmt"
	"net/http"

	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	"github.com/pedrofernandezmz/Arq-Software2/services/repositories"
	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
)

type ServiceImpl struct {
	localCache repositories.Repository
	distCache  repositories.Repository
	db         repositories.Repository
	solr       *repositories.SolrClient
}

func NewServiceImpl(
	localCache repositories.Repository,
	distCache repositories.Repository,
	db repositories.Repository,
	solr *repositories.SolrClient,
) *ServiceImpl {
	return &ServiceImpl{
		localCache: localCache,
		distCache:  distCache,
		db:         db,
		solr:       solr,
	}
}

func (serv *ServiceImpl) Get(id string) (dtos.ItemDTO, e.ApiError) {
	var item dtos.ItemDTO
	var apiErr e.ApiError
	var source string

	// try to find it in localCache
	item, apiErr = serv.localCache.Get(id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDTO{}, apiErr
		}
		// try to find it in distCache
		item, apiErr = serv.distCache.Get(id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDTO{}, apiErr
			}
			// try to find it in db
			item, apiErr = serv.db.Get(id)
			if apiErr != nil {
				if apiErr.Status() != http.StatusNotFound {
					return dtos.ItemDTO{}, apiErr
				} else {
					fmt.Println(fmt.Sprintf("Not found item %s in any datasource", id))
					apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
					return dtos.ItemDTO{}, apiErr
				}
			} else {
				source = "db"
				defer func() {
					if _, apiErr := serv.distCache.Insert(item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in distCache %v", apiErr))
					}
					if _, apiErr := serv.localCache.Insert(item); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
					}
				}()
			}
		} else {
			source = "distCache"
			defer func() {
				if _, apiErr := serv.localCache.Insert(item); apiErr != nil {
					fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Println(fmt.Sprintf("Obtained item from %s!", source))
	return item, nil
}

func (serv *ServiceImpl) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	result, apiErr := serv.db.Insert(item)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in db: %v", apiErr))
		return dtos.ItemDTO{}, apiErr
	}
	fmt.Println(fmt.Sprintf("Inserted item in db: %v", result))

	_, apiErr = serv.distCache.Insert(result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in distCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in distCache: %v", result))

	_, apiErr = serv.localCache.Insert(result)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in localCache: %v", apiErr))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in localCache: %v", result))

	apiErr2 := serv.solr.Update()
	if apiErr2 != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in solr: %v", apiErr2))
		return result, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in solr: %v", result))

	return result, nil
}
