package controller

import (
	"net/http"
	"strconv"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func GetAllRequests(repository *repository.Repository, c *gin.Context) {

	requests, err := repository.GetAllRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, requests)
}

func GetRequestByID(repository *repository.Repository, c *gin.Context) {
	var request *ds.Request

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	request, err = repository.GetRequestByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, request)
}

func DeleteRequest(repository *repository.Repository, c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	err = repository.DeleteRequest(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "deleted successful")
}

func UpdateRequest(repository *repository.Repository, c *gin.Context) {
	// Извлекаем id request из параметра запроса
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	// Попробуем извлечь JSON-данные из тела запроса
	var updatedRequest ds.Request
	if err := c.ShouldBindJSON(&updatedRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}

	err = repository.UpdateRequest(id, updatedRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func UpdateRequestStatus(repository *repository.Repository, c *gin.Context) {
	// Извлекаем id консультации из параметра запроса
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// Проверяем, что id неотрицательный
	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	// Попробуем извлечь JSON-данные из тела запроса - новый статус
	var status string
	if err := c.ShouldBindJSON(&gin.H{"status": &status}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные статуса консультации",
		})
		return
	}

	err = repository.UpdateRequestStatus(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}
