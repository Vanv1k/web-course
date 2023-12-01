package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/gin-gonic/gin"
)

// @Summary Get Consultation by ID
// @Description Show consultation by ID
// @Tags Consultations
// @ID get-consultation-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID консультации"
// @Success 200 {object} ds.Consultation
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/{id} [get]
func (c *Controller) GetConsultationByID(gctx *gin.Context) {
	var consultation *ds.Consultation

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

	consultation, err = c.Repo.GetConsultationByID(uint(id))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, consultation)
}

type Info struct {
	Name  string
	Price int
}

// @Summary Get Consultation by request ID
// @Security ApiKeyAuth
// @Description Show consultation by ID of request
// @Tags Consultations
// @ID get-consultation-by-id-of-request
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID заявки"
// @Success 200 {object} Info
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/request/{id} [get]
func (c *Controller) GetConsultationsByRequestID(gctx *gin.Context) {
	var consultationInfo ds.ConsultationInfo

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

	consultationInfo, err = c.Repo.GetConsultationsByRequestID(id)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var result []Info
	for i, _ := range consultationInfo.Names {
		consultation := Info{
			Name:  consultationInfo.Names[i],
			Price: consultationInfo.Prices[i],
		}
		result = append(result, consultation)
	}
	gctx.JSON(http.StatusOK, result)
}

// @Summary Get Consultations
// @Description Get all consultations
// @Tags Consultations
// @ID get-consultations
// @Produce json
// @Success 200 {object} ds.Consultation
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations [get]
func (c *Controller) GetAllConsultations(gctx *gin.Context) {
	maxPriceStr := gctx.DefaultQuery("maxPrice", "")
	var consultations []ds.Consultation
	var err error
	var userRequestId uint
	var maxPrice int

	if maxPriceStr != "" {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, err)
			return
		}

		consultations, userRequestId, err = c.Repo.GetConsultationsByPrice(maxPrice)
		if err != nil {
			gctx.JSON(http.StatusInternalServerError, err)
			return
		}

		gctx.JSON(http.StatusOK, gin.H{
			"consultation":    consultations,
			"ActiveRequestId": userRequestId,
		})
		return
	}

	consultations, userRequestId, err = c.Repo.GetAllConsultations()
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"consultation":    consultations,
		"ActiveRequestId": userRequestId,
	})
}

// @Summary Delete consultation by ID
// @Security ApiKeyAuth
// @Description Delete consultation by ID
// @Tags Consultations
// @ID delete-consultation-by-id
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID консультации"
// @Success 200 {string} string
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/delete/{id} [delete]
func (c *Controller) DeleteConsultation(gctx *gin.Context) {

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

	err = c.Repo.DeleteConsultation(id)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, err)
		return
	}
	gctx.JSON(http.StatusOK, "deleted successful")
}

// @Summary create consultation
// @Security ApiKeyAuth
// @Description create consultation
// @Tags Consultations
// @ID create-consultation
// @Accept json
// @Produce json
// @Param input body ds.Consultation true "consultation info"
// @Success 200 {string} string
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/create [post]
func (c *Controller) CreateConsultation(gctx *gin.Context) {
	var consultation ds.Consultation

	// Попробуйте извлечь JSON-данные из тела запроса и привести их к структуре Consultation
	if err := gctx.ShouldBindJSON(&consultation); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}

	err := c.Repo.CreateConsultation(consultation)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"consultation": consultation,
		"status":       "added",
	})
}

// @Summary update consultation
// @Security ApiKeyAuth
// @Description update consultation
// @Tags Consultations
// @ID update-consultation
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID консультации"
// @Param input body ds.Consultation true "consultation info"
// @Success 200 {string} string
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/update/{id} [put]
func (c *Controller) UpdateConsultation(gctx *gin.Context) {
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

	// Попробуем извлечь JSON-данные из тела запроса и привести их к структуре Consultation
	var updatedConsultation ds.Consultation
	if err := gctx.ShouldBindJSON(&updatedConsultation); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}
	fmt.Println(updatedConsultation)
	// Обновляем консультацию в репозитории
	err = c.Repo.UpdateConsultation(id, updatedConsultation)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

// @Summary add consultation to request
// @Security ApiKeyAuth
// @Description add consultation to request
// @Tags Consultations
// @ID add-consultation-to-request
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID консультации"
// @Success 200 {string} string
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/{id}/add-to-request [post]
func (c *Controller) AddConsultationToRequest(gctx *gin.Context) {

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

	err = c.Repo.AddConsultationToRequest(id, 1)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	gctx.JSON(http.StatusOK, gin.H{
		"status": "added to request",
	})
}

// @Summary Add consultation image
// @Security ApiKeyAuth
// @Description Add an image to a specific consultation by ID.
// @Tags Consultations
// @ID add-consultation-image
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID консультации"
// @Param image formData file true "Image file to be uploaded"
// @Success 200 {string} string
// @Failure 400 {object} ds.Consultation "Некорректный запрос"
// @Failure 404 {object} ds.Consultation "Некорректный запрос"
// @Failure 500 {object} ds.Consultation "Ошибка сервера"
// @Router /consultations/{id}/addImage [post]
func (c *Controller) AddConsultationImage(gctx *gin.Context) {
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
	// Чтение изображения из запроса
	image, err := gctx.FormFile("image")
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
		return
	}

	// Чтение содержимого изображения в байтах
	file, err := image.Open()
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при открытии"})
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения"})
		return
	}
	// Получение Content-Type из заголовков запроса
	contentType := image.Header.Get("Content-Type")

	// Вызов функции репозитория для добавления изображения
	err = c.Repo.AddConsultationImage(id, imageBytes, contentType)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	gctx.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})

}
