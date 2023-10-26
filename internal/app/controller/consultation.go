package controller

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func GetConsultationByID(repository *repository.Repository, c *gin.Context) {
	var consultation *ds.Consultation

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

	consultation, err = repository.GetConsultationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, consultation)
}

func GetConsultationsByRequestID(repository *repository.Repository, c *gin.Context) {
	var consultationInfo ds.ConsultationInfo

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

	consultationInfo, err = repository.GetConsultationsByRequestID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var result []string
	for i, _ := range consultationInfo.Names {
		result = append(result, fmt.Sprintf("%s, цена: %d рублей", consultationInfo.Names[i], consultationInfo.Prices[i]))
	}
	c.JSON(http.StatusOK, result)
}

func GetAllConsultations(repository *repository.Repository, c *gin.Context) {
	maxPriceStr := c.DefaultQuery("maxPrice", "")
	var consultations []ds.Consultation
	var err error
	var userRequestId uint
	var maxPrice int

	if maxPriceStr != "" {
		maxPrice, err = strconv.Atoi(maxPriceStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		consultations, userRequestId, err = repository.GetConsultationsByPrice(maxPrice)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, consultations)
		return
	}

	consultations, userRequestId, err = repository.GetAllConsultations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"consultation":    consultations,
		"ActiveRequestId": userRequestId,
	})
}

func DeleteConsultation(repository *repository.Repository, c *gin.Context) {

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

	err = repository.DeleteConsultation(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "deleted successful")
}

func CreateConsultation(repository *repository.Repository, c *gin.Context) {
	var consultation ds.Consultation

	// Попробуйте извлечь JSON-данные из тела запроса и привести их к структуре Consultation
	if err := c.ShouldBindJSON(&consultation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}

	err := repository.CreateConsultation(consultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"consultation": consultation,
		"status":       "added",
	})
}

func UpdateConsultation(repository *repository.Repository, c *gin.Context) {
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

	// Попробуем извлечь JSON-данные из тела запроса и привести их к структуре Consultation
	var updatedConsultation ds.Consultation
	if err := c.ShouldBindJSON(&updatedConsultation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверные данные консультации",
		})
		return
	}
	fmt.Println(updatedConsultation)
	// Обновляем консультацию в репозитории
	err = repository.UpdateConsultation(id, updatedConsultation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "updated",
	})
}

func AddConsultationToRequest(repository *repository.Repository, c *gin.Context) {

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

	err = repository.AddConsultationToRequest(id, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "added to request",
	})
}

func AddConsultationImage(repository *repository.Repository, c *gin.Context) {
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
	// Чтение изображения из запроса
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
		return
	}

	// Чтение содержимого изображения в байтах
	file, err := image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при открытии"})
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения"})
		return
	}
	// Получение Content-Type из заголовков запроса
	contentType := image.Header.Get("Content-Type")

	// Вызов функции репозитория для добавления изображения
	err = repository.AddConsultationImage(id, imageBytes, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})

}
