package product

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrapper/storage"
	"scrapper/utils"
)

type Product struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Category string    `db:"category"`
	Brand    string    `db:"brand"`
	Name     string    `db:"name"`
	Price    string    `db:"price"`
	OldPrice string    `db:"old_price"`
	//todo: add hash
}

func New(category, brand, name, price string, images []string) *Product {
	return &Product{
		Category: category,
		Brand:    brand,
		Name:     name,
		Price:    price,
	}
}

func NewFromDOMSlice(str []string, category string) *Product {
	if len(str) == 0 {
		return nil
	}

	brand := str[1]
	name := str[2]
	oldPrice, newPrice := utils.FindPriceIndexes(str, "was", "now")

	return &Product{
		Category: category,
		Brand:    brand,
		Name:     name,
		Price:    newPrice,
		OldPrice: oldPrice,
	}
}

func (p *Product) Save(ctx context.Context) error {
	if p == nil {
		return nil
	}
	res := storage.Get().WithContext(ctx).Create(p)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
