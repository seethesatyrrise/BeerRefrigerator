package rest

import (
	"BeerRefrigerator/pkg/refrigerator"
	"time"
)

type beerVM struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	ABV       string `json:"abv"`
	ExpiresAt string `json:"expiresAt"`
}

func beerDomainToVM(d *refrigerator.Beer) beerVM {
	return beerVM{
		ID:        d.ID,
		Title:     d.Title,
		ABV:       d.ABV,
		ExpiresAt: d.ExpiresAt.Format(time.RFC3339),
	}
}
