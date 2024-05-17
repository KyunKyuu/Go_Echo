package main

import (
	"log"
	"net/http"
	"os"
	"rest_gorm/app"
	"rest_gorm/controller"
	"rest_gorm/helper"
	"rest_gorm/model"
	"rest_gorm/repository"
	"rest_gorm/service"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := app.DBConnection()
	r := echo.New()

	tokenUseCase := helper.NewTokenUseCase()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, tokenUseCase)
	userController := controller.NewUserController(userService)

	addresRepo := repository.NewAddressRepo(db)
	addresService := service.NewAddressService(addresRepo, userRepo, tokenUseCase)
	addressController := controller.NewAddressController(addresService)

	r.Debug = true
	r.Validator = &CustomValidator{validator: validator.New()}
	r.HTTPErrorHandler = helper.BindAndValidate

	r.POST("/register", userController.SaveUser, JWTProtection())
	r.GET("/user/:id", userController.GetUser)
	r.GET("/users", userController.GetUserList)
	r.PUT("/user/:id", userController.UpdateUser)
	r.DELETE("/user/:id", userController.DeleteUser)
	r.GET("/user/deleted/:id", userController.GetUserDeleted)

	r.POST("/address/register", addressController.Create, JWTProtection())
	r.GET("/user/address", addressController.GetAddress, JWTProtection())
	r.GET("/address", addressController.GetAllAddress)
	r.GET("/address/:id", addressController.GetDetailAddress)
	r.PUT("/address/:id", addressController.UpdateAddress)
	r.DELETE("/address/:id", addressController.DeleteAddress)

	r.POST("/login", userController.LoginUser)
	r.Logger.Fatal(r.Start(":8080"))

}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, "anda harus login", nil))
		},
	})
}
