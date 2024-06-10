package handlers

import (
	database "customer-service/databases"
	"customer-service/models/entity"
	"customer-service/models/request"
	"customer-service/utils"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func GetUserById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	err := database.DB.Db.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func Register(c *fiber.Ctx) error {
	user := new(request.UserRequest)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: user.Password,
	}

	hashPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashPassword

	errRegister := database.DB.Db.Create(&newUser).Error
	if errRegister != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to register user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "sucess",
		"data":    newUser,
	})
}

func Login(c *fiber.Ctx) error {
	login := new(request.LoginRequest)
	if err := c.BodyParser(login); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(login)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err := database.DB.Db.First(&user, "email = ?", login.Email).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	isValid := utils.CheckPassword(login.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
