package controllers

import (
	"net/http"

	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	service "github.com/aaraya0/arq-software/arq-sw-2/services"
	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	item, apiErr := ctrl.service.Get(c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, item)
}

func (ctrl *Controller) Insert(c *gin.Context) {
	var item dtos.ItemDTO
	if err := c.BindJSON(&item); err != nil {
		apiErr := e.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	item, apiErr := ctrl.service.Insert(item)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, item)
}
