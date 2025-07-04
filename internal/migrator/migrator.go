package migrator

import "github.com/arjunksofficial/kart-challenge/internal/entities/products/models"

var (
	SampleProducts = []models.Product{
		{
			ProductMeta: models.ProductMeta{
				ID:       "1",
				Name:     "Waffle with Berries",
				Category: "Waffle",
				Price:    6.5,
			},
			Images: models.ProductImages{
				ProductID: "1",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "2",
				Name:     "Vanilla Bean Crème Brûlée",
				Category: "Crème Brûlée",
				Price:    7,
			},
			Images: models.ProductImages{
				ProductID: "2",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "3",
				Name:     "Macaron Mix of Five",
				Category: "Macaron",
				Price:    8,
			},
			Images: models.ProductImages{
				ProductID: "3",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-macaron-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-macaron-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-macaron-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-macaron-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "4",
				Name:     "Classic Tiramisu",
				Category: "Tiramisu",
				Price:    5.5,
			},
			Images: models.ProductImages{
				ProductID: "4",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-tiramisu-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-tiramisu-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "5",
				Name:     "Pistachio Baklava",
				Category: "Baklava",
				Price:    4,
			},
			Images: models.ProductImages{
				ProductID: "5",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-baklava-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-baklava-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-baklava-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-baklava-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "6",
				Name:     "Lemon Meringue Pie",
				Category: "Pie",
				Price:    5,
			},
			Images: models.ProductImages{
				ProductID: "6",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-meringue-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-meringue-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-meringue-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-meringue-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "7",
				Name:     "Red Velvet Cake",
				Category: "Cake",
				Price:    4.5,
			},
			Images: models.ProductImages{
				ProductID: "7",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-cake-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-cake-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-cake-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-cake-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "8",
				Name:     "Salted Caramel Brownie",
				Category: "Brownie",
				Price:    4.5,
			},
			Images: models.ProductImages{
				ProductID: "8",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-brownie-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-brownie-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-brownie-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-brownie-desktop.jpg",
			},
		},
		{
			ProductMeta: models.ProductMeta{
				ID:       "9",
				Name:     "Vanilla Panna Cotta",
				Category: "Panna Cotta",
				Price:    6.5,
			},
			Images: models.ProductImages{
				ProductID: "9",
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-desktop.jpg",
			},
		},
	}
)
