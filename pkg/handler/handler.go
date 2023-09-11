package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("api")
	{
		news := api.Group("/news")
		{
			news.POST("/", h.createNews)
			news.GET("/", h.getAllNews)
			news.GET("/:id", h.getNewsById)
			news.PUT("/:id", h.updateNews)
			news.DELETE("/:id", h.deleteNews)
		}
	}

	return router
}
