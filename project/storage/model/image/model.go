package image

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scrapper/storage"
)

type Image struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductId string    `db:"product_id"`
	Src       string    `db:"src"`
}

func NewImage(productId, src string) *Image {
	return &Image{
		ProductId: productId,
		Src:       src,
	}
}

func (i *Image) Save(ctx context.Context) error {
	res := storage.Get().WithContext(ctx).Create(i)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
