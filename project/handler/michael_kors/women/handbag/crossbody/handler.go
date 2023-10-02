package crossbody

import (
	"context"
	"github.com/gocolly/colly"
	"log"
	"scrapper/storage/model/image"
	"scrapper/storage/model/product"
	"scrapper/utils"
	"time"
)

const Selector = ".product-tile"

func Handle(e *colly.HTMLElement) {
	const category = "crossbody-bag"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	prodSlice := utils.GetProductString(e.Text)
	prod := product.NewFromDOMSlice(prodSlice, category)

	if prod == nil {
		return
	}

	//exists := product.NewRepository().HasByNameAndPrice(ctx, prod.Name, prod.Price)
	//if exists {
	//	log.Println("exist", prod.Name, prod.Price)
	//	return
	//}

	err := prod.Save(ctx)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	e.ForEach("img.tile-image", func(_ int, ce *colly.HTMLElement) {
		img := image.NewImage(prod.Id.String(), ce.Attr("data-src"))
		err = img.Save(ctx)
		if err != nil {
			log.Println("error on save image", err)
			return
		}
	})

	return
}
