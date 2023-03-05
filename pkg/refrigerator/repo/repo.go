package repo

import (
	"BeerRefrigerator/pkg/refrigerator"
	"context"
	"fmt"
	"github.com/go-pg/pg"
)

type repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *repo {
	return &repo{db}
}

func (r *repo) GetBeerByTitle(ctx context.Context, title string) (*refrigerator.Beer, error) {
	var res beer
	queryRes, err := r.db.QueryContext(ctx, &res,
		`SELECT * FROM beer
			   WHERE title = ?`,
		title)

	if queryRes.RowsReturned() == 0 {
		return nil, fmt.Errorf("beer with title \"%s\" not found in bd", title)
	}

	if err != nil {
		return nil, err
	}

	return beerToDomain(res), nil
}
