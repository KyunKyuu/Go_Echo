package controller

import (
	"net/http"
	"rest_gorm/model"
	"rest_gorm/model/web"
	"rest_gorm/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserControllerImpl {

	return &UserControllerImpl{
		userService: service,
	}
}

func (controller *UserControllerImpl) SaveUser(c echo.Context) error {
	user := new(web.UserServiceRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	saveUser, errSaveUser := controller.userService.SaveUser(*user)

	if errSaveUser != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errSaveUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil membuat user", saveUser))
}

func (controller *UserControllerImpl) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.userService.GetUser(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getUser))

}

func (controller *UserControllerImpl) GetUserDeleted(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	getUser, errGetUser := controller.userService.GetUserDeleted(id)

	if errGetUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errGetUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getUser))

}

func (controller *UserControllerImpl) GetUserList(c echo.Context) error {
	getUsers, errGetUsers := controller.userService.GetUserList()

	if errGetUsers != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, errGetUsers.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "success", getUsers))
}

func (controller *UserControllerImpl) UpdateUser(c echo.Context) error {
	user := new(web.UserUpdateServiceRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	userUpdate, errUserUpdate := controller.userService.UpdateUser(*user, id)

	if errUserUpdate != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errUserUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "data berhasil di update", userUpdate))
}

func (controller *UserControllerImpl) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	errDeleteUser := controller.userService.DeleteUser(id)

	if errDeleteUser != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, errDeleteUser.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "berhasil delete data", nil))

}

func (controller *UserControllerImpl) LoginUser(c echo.Context) error {
	user := new(web.UserLoginRequest)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	userLogin, errLogin := controller.userService.LoginUser(user.Email, user.Password)

	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, "login berhasil", userLogin))
}
