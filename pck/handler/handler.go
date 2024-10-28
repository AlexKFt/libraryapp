package handler

import (
	"libraryapp/pck/service"

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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		biblio := api.Group("/biblio")
		{
			biblio.POST("/", h.createBook)
			biblio.GET("/", h.getAllBooks)
			biblio.GET("/:id", h.getBookById)
			biblio.PUT("/:id", h.updateBook)
			biblio.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
