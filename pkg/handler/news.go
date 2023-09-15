package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createNews(c *gin.Context) {
	// TODO: удалить заглушку
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "test",
	})
}

func (h *Handler) getAllNews(c *gin.Context) {

}

func (h *Handler) getNewsById(c *gin.Context) {

}

func (h *Handler) updateNews(c *gin.Context) {

}

func (h *Handler) deleteNews(c *gin.Context) {

}
