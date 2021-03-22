package middleware

import (
	"fmt"
	"strings"

	"github.com/bhumong/go-user-service-v0/app/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthReq(c *fiber.Ctx) error {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return c.Next()
}

func ApiKey(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-KEY")
	if apiKey != config.Config("APP_API_KEY") {
		return fiber.NewError(400, "invalid api key")
	}
	return c.Next()
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
