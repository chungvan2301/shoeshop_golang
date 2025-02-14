package handlers

import (
	"context"
	"net/http"

	"github.com/chungvan2301/shoeshop/backend/pkg/config"
	services "github.com/chungvan2301/shoeshop/backend/pkg/repository/interface"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/response"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService services.ProductRepository
	cloudinary     *cloudinary.Cloudinary
}

func NewProductHandler(ser services.ProductRepository, config config.Config) *ProductHandler {
	cld, _ := cloudinary.NewFromURL(config.CloudinaryURL)
	return &ProductHandler{
		productService: ser,
		cloudinary:     cld,
	}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy lấy phẩm", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy sản phẩm thành công", products, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var productInput models.ProductInput

	// Bind dữ liệu từ form-data vào ProductInput (trừ ảnh)
	if err := c.ShouldBind(&productInput); err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Dữ liệu không hợp lệ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	// Lấy file ảnh từ form-data
	file, err := c.FormFile("image")
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Không có file ảnh", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// Mở file ảnh để upload lên Cloudinary
	fileData, err := file.Open()
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi mở ảnh", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	defer fileData.Close()

	// Tải ảnh lên Cloudinary
	uploadResult, err := h.cloudinary.Upload.Upload(context.Background(), fileData, uploader.UploadParams{
		Folder: "products", // Optional: Specify folder for better organization
	})
	if err != nil {
		errResponse := response.ClientResponse(http.StatusInternalServerError, "Lỗi upload ảnh lên Cloudinary", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Gán URL ảnh từ Cloudinary cho ProductInput
	productInput.ImageURL = uploadResult.SecureURL

	// Thêm sản phẩm vào database
	if err := h.productService.AddProduct(productInput); err != nil {
		errResponse := response.ClientResponse(http.StatusInternalServerError, "Lỗi thêm sản phẩm vào database", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	// Phản hồi thành công
	successResponse := response.ClientResponse(http.StatusOK, "Thêm sản phẩm thành công", productInput, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) GetCategoriesProduct(c *gin.Context) {
	categories, err := h.productService.GetCategoriesProduct()
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy category thành công", categories, nil)
	c.JSON(http.StatusOK, successResponse)
}
func (h *ProductHandler) GetGendersProduct(c *gin.Context) {
	genders, err := h.productService.GetGendersProduct()
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy category thành công", genders, nil)
	c.JSON(http.StatusOK, successResponse)
}
func (h *ProductHandler) GetBrandsProduct(c *gin.Context) {
	brands, err := h.productService.GetBrandsProduct()
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy category thành công", brands, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) GetProductsByGender(c *gin.Context) {
	gender := c.Param("gender")
	products, err := h.productService.GetProductsByGender(gender, 6)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy sản phẩm", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy sản phẩm thành công", products, nil)
	c.JSON(http.StatusOK, successResponse)
}

func (h *ProductHandler) GetProductsByBrand(c *gin.Context) {
	brand := c.Param("brand")
	products, err := h.productService.GetProductsByBrand(brand)
	if err != nil {
		errResponse := response.ClientResponse(http.StatusBadRequest, "Lỗi khi lấy sản phẩm", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "Lấy sản phẩm thành công", products, nil)
	c.JSON(http.StatusOK, successResponse)
}
