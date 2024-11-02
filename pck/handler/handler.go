package handler

import (
	"libraryapp/pck/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "libraryapp/docs"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
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
