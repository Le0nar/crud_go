package handler

import (
	"net/http"
	"strconv"

	news "github.com/Le0nar/crud_go"
	"github.com/gin-gonic/gin"
)

type createNewsInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserId int    `json:"userId" binding:"required"`
}

// @Summary Create news
// @Security ApiKeyAuth
// @Tags news
// @Description create news
// @ID create-news
// @Accept  json
// @Produce  json
// @Param input body createNewsInput true "news info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news [post]
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

type getAllNewsResponse struct {
	Data []news.News `json:"data"`
}


// @Summary Get All News 
// @Security ApiKeyAuth
// @Tags news
// @Description get all news
// @ID get-all-news
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllNewsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news [get]
func (h *Handler) getAllNews(c *gin.Context) {
	newsList, err := h.services.News.GetAllNews()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newsList)
}

// @Summary Get News By Id
// @Security ApiKeyAuth
// @Tags news
// @Description get news by id
// @ID get-news-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} news.News
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/:id [get]
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

// @Summary Update News By Id
// @Security ApiKeyAuth
// @Tags news
// @Description update news by id
// @ID update-news-by-id
// @Accept  json
// @Produce  json
// TODO: mb change string to antoher type
// @Success      200              {string}  string   "resource updated successfully"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/:id [put]
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

// @Summary Delete News By Id
// @Security ApiKeyAuth
// @Tags news
// @Description delete news by id
// @ID delete-news-by-id
// @Accept  json
// @Produce  json
// @Success      200              {string}  string    "resource deleted successfully"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/news/:id [delete]
func (h *Handler) deleteNews(c *gin.Context) {
	newsId, paramError := strconv.Atoi(c.Param("id"))

	if paramError != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err := h.services.News.DeleteNewsById(newsId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, "resource deleted successfully")
}
