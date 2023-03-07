package service

import (
	"BeerRefrigerator/pkg/refrigerator"
	"context"
)

type service struct {
	repo refrigerator.Repo
}

func NewService(repo refrigerator.Repo) *service {
	return &service{repo}
}

func (s *service) GetBeerByTitle(ctx context.Context, title string) (*refrigerator.Beer, error) {
	return s.repo.GetBeerByTitle(ctx, title)
}

func (s *service) InsertBeer(ctx context.Context, beer *refrigerator.Beer) (int64, error) {
	return s.repo.InsertBeer(ctx, beer)
}
