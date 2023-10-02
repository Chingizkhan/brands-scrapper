package product

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"scrapper/storage"
)

type Repository struct {
	conn *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{conn: storage.Get()}
}

func (r *Repository) GetByName(ctx context.Context, name string) (*Product, error) {
	p := &Product{}
	res := r.conn.WithContext(ctx).Take(p, "name = ?", name)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &Product{}, res.Error
	}
	return p, nil
}

func (r *Repository) HasByNameAndPrice(ctx context.Context, name, price string) bool {
	var exist bool
	query := "select exists(select 1 from products where name = $1 and price = $2);"
	err := r.conn.WithContext(ctx).Raw(query, name, price).Row().Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}
