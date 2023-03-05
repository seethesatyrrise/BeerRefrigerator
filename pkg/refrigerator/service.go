package refrigerator

import "context"

type Service interface {
	GetBeerByTitle(ctx context.Context, title string) (*Beer, error)
}
