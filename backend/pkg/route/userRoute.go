package route

import (
	"github.com/chungvan2301/shoeshop/backend/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup,
	productHandler *handlers.ProductHandler,
	userHandler *handlers.UserHandler,
) {

	home := router.Group("/")
	{
		home.GET("/all-products", productHandler.GetAllProducts, productHandler.GetCategoriesProduct)
		home.POST("/add-product", productHandler.AddProduct)
		home.GET("/all-categories", productHandler.GetCategoriesProduct)
		home.GET("/all-brands", productHandler.GetBrandsProduct)
		home.GET("/all-genders", productHandler.GetGendersProduct)
		home.GET("/product-list/:gender", productHandler.GetProductsByGender)
		home.GET("/:brand", productHandler.GetProductsByBrand)
	}
	user := router.Group("/user")
	{
		user.GET("/", userHandler.GetUserDetail)
		user.POST("/register-user", userHandler.RegisterUser)
		user.PUT("/edit-user", userHandler.EditUser)
		user.DELETE("/delete-user", userHandler.DeleteUser)
		user.POST("/login", userHandler.Login)
	}
}
