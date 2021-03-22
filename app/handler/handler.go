package handler

import (
	"fmt"

	"github.com/bhumong/lemonilo/app/model"
	"github.com/bhumong/lemonilo/app/repository"
	"github.com/bhumong/lemonilo/app/transformer"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repository.FindUserById(id)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "data not found",
		})
		return nil
	}
	fmt.Println(user)
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "success",
		"data":    transformer.TransformUser(user),
	})
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	users, err := repository.FindAll()
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "data not found",
		})
		return nil
	}
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "success",
		"data":    transformer.TransformUsers(users),
	})
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := repository.DeleteUserById(id)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "failed to delete user",
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
	})
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	pu := new(model.UserRequest)

	if err := c.BodyParser(pu); err != nil {
		fmt.Println(err)
		return err
	}

	id := c.Params("id")
	err := repository.UpdateUserById(id, *pu)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "failed to update user",
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
	})
	return nil
}

func CreateUser(c *fiber.Ctx) error {
	pu := new(model.UserRequest)

	if err := c.BodyParser(pu); err != nil {
		fmt.Println(err)
		return err
	}

	err := repository.CreateUser(*pu)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "failed to create user",
		})
		return nil
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	pu := new(model.UserRequest)

	if err := c.BodyParser(pu); err != nil {
		fmt.Println(err)
		return err
	}
	user, err := repository.FindUserByEmail(pu.Email)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}
	err = user.VerifyPassword(pu.Password)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "wrong password",
		})
		return nil
	}
	token, err := user.CreateToken()
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return nil
	}
	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data": &fiber.Map{
			"token": token,
		},
	})
	return nil
}
