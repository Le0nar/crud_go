package handler

import (
	"net/http"

	news "github.com/Le0nar/crud_go"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createNews(c *gin.Context) {
	var input news.News

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.News.CreateNews(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllNews(c *gin.Context) {
	newsList, err := h.services.News.GetAllNews()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newsList)
}

func (h *Handler) getNewsById(c *gin.Context) {

}

func (h *Handler) updateNews(c *gin.Context) {

}

func (h *Handler) deleteNews(c *gin.Context) {

}
