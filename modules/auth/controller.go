package auth

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/awesome/db"
	"github.com/m3rashid/awesome/models"
)

func Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		loginBody := struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}{}
		err := ctx.BodyParser(&loginBody)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}

		validate := validator.New()
		err = validate.Struct(loginBody)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Validation Error",
			})
		}

		var user models.User
		db := db.GetDb()
		err = db.Where("email = ?", loginBody.Email).First(&user).Error
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not find user",
			})
		}

		passwordsMatched := VerifyPassword(user.Password, loginBody.Password)
		if !passwordsMatched {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Credentials did not match",
			})
		}

		token, err := GenerateJWT(user.ID, user.Email)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not generate token",
			})
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"user":  user,
			"token": token,
		})
	}
}

func Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		newUser := models.User{}
		err := ctx.BodyParser(&newUser)
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}

		newUser.Deactivated = false
		newUser.Deleted = false

		validator := validator.New()
		err = validator.Struct(newUser)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}

		password := HashPassword(newUser.Password)
		newUser.Password = password

		db := db.GetDb()
		err = db.Create(&newUser).Error
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could Not Register User, Please try again later",
			})
		}

		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "User Registered Successfully",
		})
	}
}

func GetInitialUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Locals("authorized") == nil || ctx.Locals("userId") == nil || ctx.Locals("email") == nil {
			return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var user models.User
		db := db.GetDb()
		err := db.Where("id = ?", ctx.Locals("userId")).First(&user).Error
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not find user",
			})
		}

		token, err := GenerateJWT(user.ID, user.Email)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Could not generate token",
			})
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"user":  user,
			"token": token,
		})
	}
}
