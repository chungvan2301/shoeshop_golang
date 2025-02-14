package interfaces

import (
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.ProductResponse, error)
	AddProduct(product models.ProductInput) error
	GetCategoriesProduct() ([]string, error)
	GetGendersProduct() ([]string, error)
	GetBrandsProduct() ([]string, error)
	GetProductsByGender(gender string, limit int64) ([]models.ProductResponse, error)
	GetProductsByBrand(brand string) ([]models.ProductResponse, error)
}
