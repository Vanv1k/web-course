package app

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vanv1k/web-course/internal/app/ds"
	"github.com/Vanv1k/web-course/internal/app/repository"
	"github.com/Vanv1k/web-course/internal/app/role"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type registerReq struct {
	Name        string `json:"name"`
	Login       string `json:"login"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"pass"`
}

type registerResp struct {
	Ok bool `json:"ok"`
}

func (a *Application) Register(repository *repository.Repository, c *gin.Context) {
	req := &registerReq{}

	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.Password == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("password is empty"))
		return
	}

	if req.Login == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("login is empty"))
		return
	}

	if req.Name == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("name is empty"))
		return
	}

	if req.Email == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("email is empty"))
		return
	}

	if req.PhoneNumber == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("phone number is empty"))
		return
	}

	err = repository.Register(&ds.User{
		Role:        role.Buyer,
		Name:        req.Name,
		Login:       req.Login,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    generateHashString(req.Password), // пароли делаем в хешированном виде и далее будем сравнивать хеши, чтобы их не угнали с базой вместе
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &registerResp{
		Ok: true,
	})
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (a *Application) Logout(repository *repository.Repository, c *gin.Context) {
	repository.Logout()
}

type loginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
}

func (a *Application) Login(gCtx *gin.Context) {
	cfg := a.config
	req := &loginReq{}

	err := json.NewDecoder(gCtx.Request.Body).Decode(req)
	if err != nil {
		gCtx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := a.repository.GetUserByLogin(req.Login)
	fmt.Println(user)
	if err != nil {
		gCtx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(generateHashString(req.Password))
	if req.Login == user.Login && user.Password == generateHashString(req.Password) {
		// значит проверка пройдена
		// генерируем ему jwt
		cfg.JWT.SigningMethod = jwt.SigningMethodHS256
		cfg.JWT.ExpiresIn = time.Hour
		token := jwt.NewWithClaims(cfg.JWT.SigningMethod, &ds.JWTClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(cfg.JWT.ExpiresIn).Unix(),
				IssuedAt:  time.Now().Unix(),
				Issuer:    "bitop-admin",
			},
			UserID: user.Id, // test uuid
			Role:   user.Role,
		})
		if token == nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token is nil"))
			return
		}

		strToken, err := token.SignedString([]byte(cfg.JWT.Token))
		if err != nil {
			gCtx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("cant create str token"))
			return
		}

		gCtx.JSON(http.StatusOK, loginResp{
			ExpiresIn:   cfg.JWT.ExpiresIn,
			AccessToken: strToken,
			TokenType:   "Bearer",
		})
		return
	}
	fmt.Println("Response 1:")
	gCtx.AbortWithStatus(http.StatusForbidden) // отдаем 403 ответ в знак того что доступ запрещен
}
