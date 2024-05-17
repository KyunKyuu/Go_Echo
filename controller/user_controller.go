package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	SaveUser(c echo.Context) error
	GetUser(c echo.Context) error
	GetUserDeleted(c echo.Context) error
	GetUserList(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	LoginUser(c echo.Context) error
}
