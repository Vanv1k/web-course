package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Vanv1k/web-course/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const jwtPrefix = "Bearer "

func (a *Application) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(gCtx *gin.Context) {
		jwtStr := gCtx.GetHeader("Authorization")
		fmt.Println(gCtx.Request.URL.Path)
		fmt.Println(12)
		fmt.Println(strings.HasPrefix(gCtx.Request.URL.Path, "/consultations/?maxPrice"))
		fmt.Println(gCtx.Request.URL.Path == "/consultations/")

		if jwtStr == "" && (gCtx.Request.URL.Path == "/consultations/" || strings.HasPrefix(gCtx.Request.URL.Path, "/consultations/?maxPrice")) {
			fmt.Println("zdes")
		} else {
			fmt.Println("rfr")
			fmt.Println(assignedRoles)
			fmt.Println(jwtStr)
			if !strings.HasPrefix(jwtStr, jwtPrefix) { // если нет префикса то нас дурят!
				fmt.Println("ПЛОХО 1")
				gCtx.AbortWithStatus(http.StatusForbidden) // отдаем что нет доступа

				return // завершаем обработку
			}

			// отрезаем префикс
			jwtStr = jwtStr[len(jwtPrefix):]
			fmt.Println(jwtStr)

			err := a.redis.CheckJWTInBlacklist(gCtx.Request.Context(), jwtStr)
			if err == nil { // значит что токен в блеклисте
				gCtx.AbortWithStatus(http.StatusForbidden)

				return
			}
			if !errors.Is(err, redis.Nil) { // значит что это не ошибка отсуствия - внутренняя ошибка
				fmt.Println("Зашел сюда")
				fmt.Println(err)
				fmt.Println(redis.Nil)
				gCtx.AbortWithError(http.StatusInternalServerError, err)

				return
			}

			myClaims := a.ParseClaims(gCtx)

			ctxWithUserID := gCtx.Request.Context()
			ctxWithUserID = context.WithValue(ctxWithUserID, "userID", myClaims.UserID)
			gCtx.Set("userID", myClaims.UserID)

			userID, exists := gCtx.Get("userID")
			if exists {
				fmt.Println(userID.(uint))
			}

			ctxWithUserRole := gCtx.Request.Context()
			ctxWithUserRole = context.WithValue(ctxWithUserRole, "userRole", myClaims.Role)
			gCtx.Set("userRole", myClaims.Role)

			userRole, exists := gCtx.Get("userRole")
			if exists {
				fmt.Println(userRole.(role.Role))
			}

			fmt.Println("Сюда()")
			fmt.Println(myClaims)
			authorized := false

			for _, oneOfAssignedRole := range assignedRoles {
				if myClaims.Role == oneOfAssignedRole {
					authorized = true
					break
				}
			}

			if !authorized {
				gCtx.AbortWithStatus(http.StatusForbidden)
				log.Printf("role %s is not assigned in %s", myClaims.Role, assignedRoles)
				return
			}
		}

	}

}
