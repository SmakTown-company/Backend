package handler

import (
	"github.com/SmakTown-company/Backend/payment/internal/service"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	})
	router.Use(corsConfig)
	allowedCIDRs := []string{
		"127.0.0.1/32", // IPv4 localhost
		"::1/128",      // IPv6 localhost
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/CreatePayment", h.CreatePayment)
	router.POST("/webhooks", h.IPFilterMiddleware(h.PaymentStatus, allowedCIDRs))

	return router
}
