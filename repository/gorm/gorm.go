package gorm

import (
	"context"

	"gorm.io/gorm"
)

type GormBaseRepository[E any, K any] struct {
	db        *gorm.DB
	ModelName string
}

func (r *GormBaseRepository[E, K]) Model() string {
	return r.ModelName
}

func (r *GormBaseRepository[E, K]) getDB(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *GormBaseRepository[E, K]) ListAll(ctx context.Context) ([]E, error) {
	var ents []E
	res := r.getDB(ctx).Model(&ents).Find(&ents)
	return ents, res.Error
}

func (r *GormBaseRepository[E, K]) Find(ctx context.Context) ([]E, error) {
	var ents []E
	res := r.getDB(ctx).Model(&ents).Find(&ents)
	return ents, res.Error
}

func (r *GormBaseRepository[E, K]) FindByKey(ctx context.Context, key K) (E, error) {
	var ent E
	err := r.getDB(ctx).First(&ent, key).Error
	return ent, err
}

func (r *GormBaseRepository[E, K]) Update(ctx context.Context, ent *E) error {
	err := r.getDB(ctx).Save(ent).Error
	return err
}

func (r *GormBaseRepository[E, K]) Create(ctx context.Context, ent *E) error {
	err := r.getDB(ctx).Create(ent).Error
	return err
}

func (r *GormBaseRepository[E, K]) Delete(ctx context.Context, ent *E) error {
	err := r.getDB(ctx).Delete(ent).Error
	return err
}

func NewGormBaseRepository[E any, K any](
	db *gorm.DB,
	model string,

) *GormBaseRepository[E, K] {
	return &GormBaseRepository[E, K]{
		db:        db,
		ModelName: model,
	}
}

func ProvideGormBaseRepository[E any, K any](
	db *gorm.DB,
	model string,
) *GormBaseRepository[E, K] {
	return NewGormBaseRepository[E, K](db, model)
}
