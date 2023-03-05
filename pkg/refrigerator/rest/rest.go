package rest

import (
	"BeerRefrigerator/pkg/refrigerator"
	"BeerRefrigerator/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rest struct {
	service refrigerator.Service
}

func NewRest(service refrigerator.Service) *rest {
	return &rest{service}
}

func (r *rest) Register(api *gin.RouterGroup) {
	route := api.Group("/refrigerator")
	{
		route.GET(":title", r.getBeerByTitle)
	}
}

func (r *rest) getBeerByTitle(c *gin.Context) {
	ctx := c.Request.Context()
	title := c.Param("title")
	if title == "" {
		utils.PublishError(c, errors.New("empty title"), http.StatusBadRequest)
		return
	}

	beer, err := r.service.GetBeerByTitle(ctx, title)
	if err != nil {
		utils.PublishError(c, err, http.StatusInternalServerError)
		return
	}

	utils.PublishData(c, beerDomainToVM(beer))
}
