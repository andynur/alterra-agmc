package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtUtils struct{}

type JwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

type TokenData struct {
	UserId    uint      `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (u *JwtUtils) GenerateToken(userId uint) (data TokenData, err error) {
	expiredTime := time.Now().Add(time.Hour * 1).Unix() // 1 hours

	// Set custom claims
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			Issuer:    "AGMC",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("EMPTY JWT_SECRET ENV")
	}

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return data, err
	}

	return TokenData{
		UserId:    userId,
		Token:     signedToken,
		ExpiredAt: time.Unix(expiredTime, 0),
	}, nil
}

func (u *JwtUtils) ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		return int(claims["user_id"].(float64))
	}

	return 0
}
