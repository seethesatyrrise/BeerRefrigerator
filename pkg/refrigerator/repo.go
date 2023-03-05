package refrigerator

import "context"

type Repo interface {
	GetBeerByTitle(ctx context.Context, title string) (*Beer, error)
}
