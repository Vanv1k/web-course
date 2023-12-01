package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/gin-gonic/gin"
)

// @Summary Get Requests
// @Security ApiKeyAuth
// @Description Get all requests
// @Tags Requests
// @ID get-requests
// @Produce json
// @Success 200 {object} ds.Request
// @Failure 400 {object} ds.Request "Некорректный запрос"
// @Failure 404 {object} ds.Request "Некорректный запрос"
// @Failure 500 {object} ds.Request "Ошибка сервера"
// @Router /requests [get]
func (c *Controller) GetAllRequests(gctx *gin.Context) {
	status := gctx.DefaultQuery("status", "")
	startFormationDateStr := gctx.DefaultQuery("startDate", "")
	endFormationDateStr := gctx.DefaultQuery("endDate", "")
	var requests []ds.Request
	var err error

	if status != "" {
		requests, err = c.Repo.GetRequestsByStatus(status)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, err)
			return
		}

		gctx.JSON(http.StatusOK, requests)
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
			gctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if endFormationDateStr != "" {
			endFormationDate, err = time.Parse(layout, endFormationDateStr)

			if err != nil {
				gctx.JSON(http.StatusInternalServerError, err)
				return
			}
		}

		requests, err = c.Repo.GetRequestsByDate(startFormationDate, endFormationDate)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, err)
			return
		}

		gctx.JSON(http.StatusOK, requests)
		return
	}
	log.Println("go here")
	requests, err = c.Repo.GetAllRequests()
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, requests)
}

// @Summary Delete Request by ID
// @Security ApiKeyAuth
// @Description Delete request by ID
// @Tags Requests
// @ID delete-request-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID заявки"
// @Success 200 {string} string
// @Failure 400 {object} ds.Request "Некорректный запрос"
// @Failure 404 {object} ds.Request "Некорректный запрос"
// @Failure 500 {object} ds.Request "Ошибка сервера"
// @Router /requests/delete/{id} [delete]
func (c *Controller) DeleteRequest(gctx *gin.Context) {

	id, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if id < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	err = c.Repo.DeleteRequest(id)

	if err != nil {
		gctx.JSON(http.StatusBadRequest, err)
		return
	}
	gctx.JSON(http.StatusOK, "deleted successful")
}

// @Summary Update Request by ID
// @Security ApiKeyAuth
// @Description Update request by ID
// @Tags Requests
// @ID update-request-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID заявки"
// @Param input body ds.Request true "request info"
// @Success 200 {string} string
// @Failure 400 {object} ds.Request "Некорректный запрос"
// @Failure 404 {object} ds.Request "Некорректный запрос"
// @Failure 500 {object} ds.Request "Ошибка сервера"
// @Router /requests/update/{id} [put]
func (c *Controller) UpdateRequest(gctx *gin.Context) {
	// Извлекаем id request из параметра запроса
	id, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if id < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	// Попробуем извлечь JSON-данные из тела запроса
	var updatedRequest ds.Request
	if err := gctx.ShouldBindJSON(&updatedRequest); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}

	err = c.Repo.UpdateRequest(id, updatedRequest)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

// @Summary Update Request Status By Moderator
// @Security ApiKeyAuth
// @Description Update request by moderator
// @Tags Requests
// @ID update-request-status-by-moderator
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID заявки"
// @Param input body ds.StatusData true "status info"
// @Success 200 {string} string
// @Failure 400 {object} ds.Request "Некорректный запрос"
// @Failure 404 {object} ds.Request "Некорректный запрос"
// @Failure 500 {object} ds.Request "Ошибка сервера"
// @Router /requests/{id}/moderator/update-status [put]
func (c *Controller) UpdateRequestStatus(gctx *gin.Context) {
	// Извлекаем id консультации из параметра запроса
	id, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Проверяем, что id неотрицательный
	if id < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	// Попробуем извлечь JSON-данные из тела запроса - новый статус
	var status ds.StatusData
	if err := gctx.ShouldBindJSON(&status); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные статуса заявки",
		})
		return
	}
	statusStr := status.Status
	log.Println(statusStr)
	if statusStr != "canceled" && statusStr != "finished" {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные статуса заявки",
		})
		return
	}

	err = c.Repo.UpdateRequestStatus(id, statusStr)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

// @Summary Update Request Status By User
// @Security ApiKeyAuth
// @Description Update request status by user
// @Tags Requests
// @ID update-request-status-by-user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID заявки"
// @Success 200 {string} string
// @Failure 400 {object} ds.Request "Некорректный запрос"
// @Failure 404 {object} ds.Request "Некорректный запрос"
// @Failure 500 {object} ds.Request "Ошибка сервера"
// @Router /requests/{id}/user/update-status [put]
func (c *Controller) UpdateRequestStatusToSendedByUser(gctx *gin.Context) {
	// Извлекаем id заявки из параметра запроса
	id, err := strconv.Atoi(gctx.Param("id"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Проверяем, что id неотрицательный
	if id < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}
	err = c.Repo.UpdateRequestStatus(id, "formed")
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"status": "updated, new status - `formed`",
	})
}
