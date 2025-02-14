package http

import (
	"fmt"
	"time"

	"github.com/chungvan2301/shoeshop/backend/pkg/api/handlers"
	"github.com/chungvan2301/shoeshop/backend/pkg/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServer(
	productHandler *handlers.ProductHandler,
	userHandler *handlers.UserHandler,
) *ServerHTTP {

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},        // Nguồn được phép
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Phương thức được phép
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Use(gin.Logger())

	route.UserRoute(engine.Group("/"), productHandler, userHandler)

	return &ServerHTTP{
		engine: engine,
	}
}

// Start khởi chạy server
func (s *ServerHTTP) Start() {
	if s.engine == nil {
		fmt.Println("Có lỗi!")
	}
	s.engine.Run(":8080")
}
