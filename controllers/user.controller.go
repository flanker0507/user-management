package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-todo-app/database"
	"go-todo-app/models"
	"go-todo-app/request"
	"go-todo-app/utils"
	"log"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var existingUser models.User
	result := database.DB.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	newUser := models.User{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
func UpdateUserById(ctx *fiber.Ctx) error {
	userReq := request.UserUpdateRequest{}

	// PARSE REQUEST BODY
	if errParse := ctx.BodyParser(&userReq); errParse != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "fail to parse data",
			"error":   errParse.Error(),
		})
	}

	// VALIDATE REQUEST DATA
	validate := validator.New()
	if errValidate := validate.Struct(&userReq); errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "some data is not valid",
			"error":   errValidate.Error(),
		})
	}

	userId := ctx.Params("id")
	user := models.User{}

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if user.Email != userReq.Email {
		var existingUser models.User
		if err := database.DB.Where("email = ?", userReq.Email).First(&existingUser).Error; err == nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message": "email already exists",
			})
		}
	}
	user.Name = userReq.Name
	user.Email = userReq.Email

	if errSave := database.DB.Save(&user).Error; errSave != nil {
		log.Println("Error saving user:", errSave)
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user updated",
		"data":    user,
	})
}

func DeleteUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user := models.User{}

	if err := database.DB.First(&user, "id = ?", &userId).Error; err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if errDel := database.DB.Delete(&user).Error; errDel != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user deleted",
	})
}
