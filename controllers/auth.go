package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"crud.com/crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/exp/slices"
)

var jwtSecret = []byte("key")

// @Summary Sso
// @Description sso
// @Tags    user
// @Accept  json
// @Produce  json
// @Param   user   body      string  true  "sso"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /sso [post]
func (c *Controller) Sso(ctx *gin.Context) {
	var data = new(models.User)

	if ctx.ShouldBindJSON(&data) != nil {
		res := models.Result{
			IsSuccess:     false,
			ReturnCode:    models.ModelInValid,
			ReturnMessage: "ModelInValid",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	user, err := GetMockUser(data.ID)
	if err != nil {
		res := models.Result{
			IsSuccess:     false,
			ReturnCode:    err.Error(),
			ReturnMessage: "UserNotFound",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	// set claims and sign
	claims := models.Claims{
		Account: user.Username,
		Role:    user.ID,
		StandardClaims: jwt.StandardClaims{
			Audience:  user.Username,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        user.ID,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ginJWT",
			NotBefore: time.Now().Add(10 * time.Second).Unix(),
			Subject:   user.Username,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		res := models.Result{
			IsSuccess:     false,
			ReturnCode:    models.JwtTokenGenFail,
			ReturnMessage: err.Error(),
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res := models.ResultWithData[string]{
		Data: token,
	}
	res.IsSuccess = true
	res.ReturnCode = models.Success
	res.ReturnMessage = ""

	ctx.JSON(http.StatusOK, res)
	return
}

// validate JWT
func (c *Controller) AuthRequired(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	splitAuthorization := strings.SplitN(auth, "Bearer ", 2)
	if len(splitAuthorization) != 2 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	token := strings.Split(auth, "Bearer ")[1]

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		ctx.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
		fmt.Println(claims)
		fmt.Println("account:", claims.Account)
		fmt.Println("role:", claims.Role)
		ctx.Set("account", claims.Account)
		ctx.Set("role", claims.Role)
		ctx.Next()
	} else {
		ctx.Abort()
		return
	}
}

func GetMockUser(id string) (*models.User, error) {
	list := []models.User{
		{ID: "00001", Username: "admin", Phone: "0911123456", Email: "aa"},
		{ID: "00002", Username: "SA", Phone: "0922123456", Email: "bb"},
		{ID: "00003", Username: "PG", Phone: "0933123456", Email: "cc"},
	}

	result := slices.IndexFunc(list, func(u models.User) bool { return u.ID == id })
	fmt.Println(result)
	if result == -1 {
		err := errors.New(models.UserNotFound)
		return nil, err
	}

	return &list[result], nil
}
