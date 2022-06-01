package controllers

import (
	"errors"
	"fmt"
	"net/http"
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

func CheckToken() {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
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
