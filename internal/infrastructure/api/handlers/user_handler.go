package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserHandler struct {
	userService ports.IUserService
}

// SignUp es un handler para registrarse en el sistema.
//
// Recibe in SignUpRequest, que contiene toda la información del usuario, la
// guarda en la base de datos, y finalmente retorna un JSON con el ID, nombre
// apellido y el token JWT generado.
//
// Estados HTTP:
//
// 400: El SignUpRequest tiene un formato incorrecto.
// 500: Ha ocurrido un error inesperado al registrar el usuario.
// 201: El usuario ha sido creado con éxito.
func (uc *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.SignUp(*signUpRequest.ToUserCore())
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusCreated, authResponse)
}

// SignIn es un handler para iniciar sesión en el sistema.
//
// Recibe in SignInRequest, que contiene el correo y la contraseña del usuario, 
// recupera al usuario de la base de datos usando el correo, valida si la
// contraseña es correcta y finalmente retorna un JSON con el ID, nombre, apellido
// y el token JWT generado.
//
// Estados HTTP:
//
// 400: El SignInRequest tiene un formato incorrecto.
// 401: El usuario no se encuentra autorizado.
// 500: Ha ocurrido un error inesperado al registrar el usuario.
// 201: El usuario ha sido creado con éxito.
func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var authResponse dto.AuthResponse
	authResponse.LoadFromAuthorizedUserCore(*user)
	c.JSON(http.StatusOK, authResponse)
}

func (uc *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var userResponse []dto.UserResponse
	for _, k := range users {
		var res dto.UserResponse
		res.LoadFromUserCore(k)
		userResponse = append(userResponse, res)
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserHandler) FindUser(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.HTTPError(core.BadRequest, c)
		return
	}

	user, err := uc.userService.FindUserById(userID)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	var userResponse dto.UserResponse
	userResponse.LoadFromUserCore(*user)

	authResponse := dto.AuthResponse{
		User:        userResponse,
		AccessToken: "",
	}

	c.JSON(http.StatusOK, authResponse)
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}
