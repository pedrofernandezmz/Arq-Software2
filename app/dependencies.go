package app

import (
	"github.com/pedrofernandezmz/Arq-Software2/controllers"
	service "github.com/pedrofernandezmz/Arq-Software2/services"
	"github.com/pedrofernandezmz/Arq-Software2/services/repositories"
	"time"
)

type Dependencies struct {
	ItemController *controllers.Controller
}

func BuildDependencies() *Dependencies {
	// Repositories
	ccache := repositories.NewCCache(1000, 100, 30*time.Second)
	memcached := repositories.NewMemcached("localhost", 11211)
	mongo := repositories.NewMongoDB("localhost", 27017, "avisos")
	solr := repositories.NewSolrClient("localhost", 8983, "avisos")
	
	// Services
	service := service.NewServiceImpl(ccache, memcached, mongo, solr)

	// Controllers
	controller := controllers.NewController(service)

	return &Dependencies{
		ItemController: controller,
	}
}
