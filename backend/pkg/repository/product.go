package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/chungvan2301/shoeshop/backend/pkg/domain"
	"github.com/chungvan2301/shoeshop/backend/pkg/ultis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	productCollection *mongo.Collection
	ctx               context.Context
}

func NewProductRepository(productCollection *mongo.Collection, ctx context.Context) *ProductRepository {
	return &ProductRepository{
		productCollection: productCollection,
		ctx:               ctx,
	}
}

func (u *ProductRepository) GetAllProducts() ([]models.ProductResponse, error) {

	var products []models.ProductResponse
	ctx := u.ctx
	cursor, err := u.productCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Lỗi khi tìm sản phẩm:", err)
		return products, err
	}
	defer cursor.Close(u.ctx)

	for cursor.Next(ctx) {
		var product domain.Product
		if err = cursor.Decode(&product); err != nil {
			log.Println("Lỗi khi giải mã sản phẩm:", err)
			continue
		}

		productResponse := models.ProductResponse{
			ID:       product.ID.Hex(),
			Name:     product.Name,
			Brand:    product.Brand,
			Price:    product.Price,
			ImageURL: product.ImageURL,
		}

		products = append(products, productResponse)
	}

	if err = cursor.Err(); err != nil {
		log.Println("Lỗi cursor:", err)
	}

	return products, err
}
func (u *ProductRepository) AddProduct(product models.ProductInput) error {
	// Chuyển đổi ProductInput thành Product để lưu vào cơ sở dữ liệu

	newProduct := domain.Product{
		ID:            primitive.NewObjectID(),
		Name:          product.Name,
		Brand:         product.Brand,
		Gender:        product.Gender,
		Category:      product.Category,
		Price:         product.Price,
		IsInInventory: product.IsInInventory,
		ItemsLeft:     product.ItemsLeft,
		ImageURL:      product.ImageURL,
		Slug:          product.Slug,
	}

	// Thực hiện thao tác InsertOne để thêm vào MongoDB
	_, err := u.productCollection.InsertOne(u.ctx, newProduct)
	if err != nil {
		log.Println("Lỗi khi thêm sản phẩm:", err)
		return err
	}

	log.Println("Thêm sản phẩm thành công:", newProduct.Name)
	return nil
}

func (r *ProductRepository) GetCategoriesProduct() ([]string, error) {
	// Tạo một pipeline để lấy danh sách category, loại bỏ trùng lặp và sắp xếp từ A-Z
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$category"}}}}, // Group theo category
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},            // Sắp xếp từ A-Z
	}

	cursor, err := r.productCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("error aggregating categories: %v", err)
	}
	defer cursor.Close(context.Background())

	var categories []string
	for cursor.Next(context.Background()) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding category: %v", err)
		}
		categories = append(categories, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return categories, nil
}

func (r *ProductRepository) GetGendersProduct() ([]string, error) {
	// Tạo một pipeline để lấy danh sách category, loại bỏ trùng lặp và sắp xếp từ A-Z
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$gender"}}}},
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
	}

	cursor, err := r.productCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("error aggregating gender: %v", err)
	}
	defer cursor.Close(context.Background())

	var genders []string
	for cursor.Next(context.Background()) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding genders: %v", err)
		}
		genders = append(genders, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return genders, nil
}

func (r *ProductRepository) GetBrandsProduct() ([]string, error) {
	// Tạo một pipeline để lấy danh sách category, loại bỏ trùng lặp và sắp xếp từ A-Z
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$brand"}}}},
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
	}

	cursor, err := r.productCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("error aggregating gender: %v", err)
	}
	defer cursor.Close(context.Background())

	var brands []string
	for cursor.Next(context.Background()) {
		var result struct {
			ID string `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("error decoding brands: %v", err)
		}
		brands = append(brands, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return brands, nil
}

func (u *ProductRepository) GetProductsByGender(gender string, limit int64) ([]models.ProductResponse, error) {
	var products []models.ProductResponse

	filter := bson.M{"gender": gender}

	cursor, err := u.productCollection.Find(u.ctx, filter)
	if err != nil {
		log.Println("Lỗi khi tìm sản phẩm theo giới tính:", err)
		return products, err
	}
	defer cursor.Close(u.ctx)

	// Chạy qua các kết quả trả về
	for cursor.Next(u.ctx) {
		var product domain.Product
		if err = cursor.Decode(&product); err != nil {
			log.Println("Lỗi khi giải mã sản phẩm:", err)
			continue
		}

		productResponse := models.ProductResponse{
			ID:       product.ID.Hex(),
			Name:     product.Name,
			Brand:    product.Brand,
			Price:    product.Price,
			ImageURL: product.ImageURL,
			Gender:   product.Gender,
		}

		products = append(products, productResponse)

		if int64(len(products)) >= limit {
			break
		}
	}

	if err = cursor.Err(); err != nil {
		log.Println("Lỗi cursor:", err)
	}

	return products, err
}

func (u *ProductRepository) GetProductsByBrand(brand string) ([]models.ProductResponse, error) {
	var products []models.ProductResponse

	filter := bson.M{"brand": brand}

	cursor, err := u.productCollection.Find(u.ctx, filter)
	if err != nil {
		log.Println("Lỗi khi tìm sản phẩm theo hãng:", err)
		return products, err
	}
	defer cursor.Close(u.ctx)

	for cursor.Next(u.ctx) {
		var product domain.Product
		if err = cursor.Decode(&product); err != nil {
			log.Println("Lỗi khi giải mã sản phẩm:", err)
			continue
		}

		productResponse := models.ProductResponse{
			ID:       product.ID.Hex(),
			Name:     product.Name,
			Brand:    product.Brand,
			Price:    product.Price,
			ImageURL: product.ImageURL,
		}

		products = append(products, productResponse)
	}

	if err = cursor.Err(); err != nil {
		log.Println("Lỗi cursor:", err)
	}

	return products, err
}
