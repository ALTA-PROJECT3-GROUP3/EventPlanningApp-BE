package helper

import (
	"time"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"

	"github.com/golang-jwt/jwt/v4"
	echoJWT "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echoJWT.WithConfig(echoJWT.Config{
		SigningKey:    []byte(config.JWTKEY),
		SigningMethod: "HS256",
	})
}

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTKEY))
}

func DecodeToken(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return uint(userId)
	}
	return 0
}
