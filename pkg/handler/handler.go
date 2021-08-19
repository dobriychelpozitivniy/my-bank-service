package handler

import (
	"github.com/dobriychelpozitivniy/mybank/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	account := router.Group("/account")
	{
		account.PATCH("/add", h.Add)
		account.GET("/balance", h.Balance)
		account.PATCH("/withdraw", h.Withdraw)
	}

	return router
}
