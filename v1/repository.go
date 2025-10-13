package repository

import "context"

type Repository[E any, K any] interface {
	Model() string
	ListAll(ctx context.Context) ([]E, error)
	Find(ctx context.Context) ([]E, error)
	FindByKey(ctx context.Context, key K) (E, error)
	Update(ctx context.Context, ent *E) error
	Create(ctx context.Context, ent *E) error
	Delete(ctx context.Context, ent *E) error
}
