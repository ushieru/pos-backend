package api_v1

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/ushieru/pos/database"
	"github.com/ushieru/pos/models"
	"github.com/ushieru/pos/models/errors"
	"github.com/ushieru/pos/utils"
)

func setupAuthRoutes(app fiber.Router) {
	session := app.Group("/auth")
	session.Post("/login", login)
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// @Router /api/v1/auth/login [POST]
// @Param credentials body Login true "credentials"
// @Tags Auth
// @Accepts json
// @Produce json
// @Success 200 {object} LoginResponse
// @Failure 0 {object} models_errors.ErrorResponse
func login(c *fiber.Ctx) error {
	loginParams := new(Login)
	if err := c.BodyParser(loginParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse("Params error", "Params error", ""))
	}
	errors := utils.ValidateStruct(loginParams)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	var account models.Account
	database.DBConnection.First(&account, "Username = ?", loginParams.Username)
	if account.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse("Auth Error", "Error en usuario o contraseña", ""))
	}
	res := utils.CheckPasswordHash(loginParams.Password, account.Password)
	if !res {
		return c.Status(fiber.StatusBadRequest).JSON(
			models_errors.NewErrorResponse("Auth Error", "Error en usuario o contraseña", ""))
	}
	claims := jwt.MapClaims{
		utils.SessionParamAdminId: account.ID,
		utils.SessionParamUserId:  account.UserID,
		utils.SessionParamRole:    string(account.AccountType),
		"exp":                     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("super secret word")) // TODO: Change "super secret word"
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var user models.User
	database.DBConnection.Preload("Account").First(&user, account.UserID)
	return c.JSON(LoginResponse{
		Token: t,
		User:  user,
	})
}
