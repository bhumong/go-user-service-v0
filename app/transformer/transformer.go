package transformer

import (
	"time"

	"github.com/bhumong/go-user-service-v0/app/model"
	"github.com/gofiber/fiber/v2"
)

func TransformUser(user model.User) *fiber.Map {
	address := ""
	if user.Address.Valid {
		address = user.Address.String
	}
	createdAt := ""
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time.Format(time.RFC3339)
	}
	updatedAt := ""
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format(time.RFC3339)
	}
	return &fiber.Map{
		"userId":    user.UserId,
		"email":     user.Email,
		"address":   address,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
	}
}

func TransformUsers(users model.Users) map[int]interface{} {
	m := make(map[int]interface{})
	for index, user := range users.Users {
		m[index] = TransformUser(user)
	}
	return m
}
