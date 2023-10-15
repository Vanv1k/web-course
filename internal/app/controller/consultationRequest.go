package controller

import (
	"net/http"
	"strconv"

	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/gin-gonic/gin"
)

func DeleteConsultationRequest(repository *repository.Repository, c *gin.Context) {
	var idC, idR int
	var err error
	idC, err = strconv.Atoi(c.Param("id_c"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if idC < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id консультации",
		})
		return
	}

	idR, err = strconv.Atoi(c.Param("id_r"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if idR < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id заявки",
		})
		return
	}

	err = repository.DeleteConsultationRequest(idC, idR)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "deleted successful")
}
