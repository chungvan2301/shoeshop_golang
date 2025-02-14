package di

import (
	"context"
	"log"

	http "github.com/chungvan2301/shoeshop/backend/pkg/api"
	"github.com/chungvan2301/shoeshop/backend/pkg/api/handlers"
	"github.com/chungvan2301/shoeshop/backend/pkg/config"
	"github.com/chungvan2301/shoeshop/backend/pkg/db"
	"github.com/chungvan2301/shoeshop/backend/pkg/repository"
)

func InitializeApp() {
	// Tải cấu hình từ file hoặc biến môi trường
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	// Kết nối MongoDB
	_, err := db.ConnectMongoDB(config)
	if err != nil {
		log.Fatalf("Lỗi kết nối Mongo: %v", err)
	}

	// Khởi tạo repository và handler
	productRepository := repository.NewProductRepository(db.GetProductCollection(), context.Background())
	productHandler := handlers.NewProductHandler(productRepository, config)

	userRepository := repository.NewUserRepository(db.GetUserCollection(), context.Background())
	userHandler := handlers.NewUserHandler(userRepository)

	// Khởi tạo server HTTP và khởi chạy
	server := http.NewServer(productHandler, userHandler)
	server.Start()
}
