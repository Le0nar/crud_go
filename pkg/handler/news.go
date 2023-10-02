package handler

import (
	"net/http"
	"strconv"

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
	itemId, paramError := strconv.Atoi(c.Param("id"))

	if paramError != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	news, err := h.services.News.GetNewsById(itemId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, news)
}

func (h *Handler) updateNews(c *gin.Context) {
	newsId, paramError := strconv.Atoi(c.Param("id"))

	if paramError != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}


	var input news.UpdateNewsInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	isTitleExist := input.Title != nil
	isDescriptionExist := input.Description != nil

	if !isTitleExist && !isDescriptionExist {
		newErrorResponse(c, http.StatusBadRequest, "invalid payload")
		return
	}

	err := h.services.UpdateNewsById(input, newsId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "resource updated successfully")
}

func (h *Handler) deleteNews(c *gin.Context) {
}
