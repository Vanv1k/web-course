package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Delete Consultation From Request
// @Security ApiKeyAuth
// @Description delete consultation from request
// @Tags Consultation-Request
// @ID delete-consultation-from-request
// @Accept       json
// @Produce      json
// @Param        id_c   path      int  true  "ID консультации"
// @Param        id_r   path      int  true  "ID заявки"
// @Success 200 {string} string "Консультация была удалена из заявки"
// @Failure 400 {string} string "Некорректный запрос"
// @Failure 404 {string} string "Некорректный запрос"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /consultation-request/delete/consultation/{id_c}/request/{id_r} [delete]
func (c *Controller) DeleteConsultationRequest(gctx *gin.Context) {
	var idC, idR int
	var err error
	idC, err = strconv.Atoi(gctx.Param("id_c"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if idC < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id консультации",
		})
		return
	}

	idR, err = strconv.Atoi(gctx.Param("id_r"))
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if idR < 0 {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "неверное значение id заявки",
		})
		return
	}

	err = c.Repo.DeleteConsultationRequest(idC, idR)

	if err != nil {
		gctx.JSON(http.StatusBadRequest, err)
		return
	}
	gctx.JSON(http.StatusOK, "deleted successful")
}
