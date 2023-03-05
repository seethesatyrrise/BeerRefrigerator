package service

import (
	refrigerator2 "BeerRefrigerator/pkg/refrigerator"
	"context"
)

type service struct {
	repo refrigerator2.Repo
}

func NewService(repo refrigerator2.Repo) *service {
	return &service{repo}
}

func (s *service) GetBeerByTitle(ctx context.Context, title string) (*refrigerator2.Beer, error) {
	return s.repo.GetBeerByTitle(ctx, title)
}
