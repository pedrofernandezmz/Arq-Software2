package app

import (
	"github.com/aaraya0/arq-software/arq-sw-2/controllers"
	service "github.com/aaraya0/arq-software/arq-sw-2/services"
	"github.com/aaraya0/arq-software/arq-sw-2/services/repositories"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	//ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("localhost", 11211)
	mongo := repositories.NewMongoDB("localhost", 27017, "publicaciones")
	solr := repositories.NewSolrClient("localhost", 8983, "publicaciones")
	// Services
	service := service.NewServiceImpl(memcached, mongo, solr)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
