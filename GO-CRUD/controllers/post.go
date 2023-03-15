package controllers

import (
	"GO-CRUD/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
N.B : binding works the same way as validator in laravel,
ain't it wonderful ?
*/
type CreatePostValidator struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostValidator
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}
