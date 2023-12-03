package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (a *Application) ParseClaims(gCtx *gin.Context) *ds.JWTClaims {

	jwtStr := gCtx.GetHeader("Authorization")
	jwtStr = jwtStr[len(jwtPrefix):] // отрезаем префикс
	fmt.Println(jwtStr)
	token, err := jwt.ParseWithClaims(jwtStr, &ds.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.JWT.Token), nil
	})
	if err != nil {
		gCtx.AbortWithStatus(http.StatusForbidden)
		log.Println(err)

		return nil
	}

	myClaims := token.Claims.(*ds.JWTClaims)
	return myClaims
}

func (a *Application) ParseUserID(gCtx *gin.Context) uint {
	myClaims := a.ParseClaims(gCtx)
	return myClaims.UserID
}
