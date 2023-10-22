package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func GetAllRequests(repository *repository.Repository, c *gin.Context) {
	status := c.DefaultQuery("status", "")
	startFormationDateStr := c.DefaultQuery("startDate", "")
	endFormationDateStr := c.DefaultQuery("endDate", "")
	var requests []ds.Request
	var err error

	if status != "" {
		requests, err = repository.GetRequestsByStatus(status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, requests)
		return
	}
	log.Println(startFormationDateStr + "ASSDA")
	if startFormationDateStr != "" {
		var startFormationDate time.Time
		var endFormationDate time.Time
		layout := "2006-01-02 15:04:05.000000"
		startFormationDate, err = time.Parse(layout, startFormationDateStr)
		log.Println(startFormationDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		if endFormationDateStr != "" {
			endFormationDate, err = time.Parse(layout, endFormationDateStr)

			if err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}
		}

		requests, err = repository.GetRequestsByDate(startFormationDate, endFormationDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, requests)
		return
	}
	log.Println("go here")
	requests, err = repository.GetAllRequests()
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
	var status ds.StatusData
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные статуса заявки",
		})
		return
	}
	statusStr := status.Status
	log.Println(statusStr)

	err = repository.UpdateRequestStatus(id, statusStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func UpdateRequestStatusToSendedByUser(repository *repository.Repository, c *gin.Context) {
	// Извлекаем id заявки из параметра запроса
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
	err = repository.UpdateRequestStatus(id, "formed")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated, new status - `formed`",
	})
}
