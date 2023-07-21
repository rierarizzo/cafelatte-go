package handlers

import (
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserHandler struct {
	userService ports.IUserService
}

func (uc *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	if err := c.BindJSON(&signUpRequest); err != nil {
		utils.HTTPError(errors.ErrBadRequest, c)
		return
	}

	user, err := uc.userService.SignUp(*mappers.FromSignUpRequestToUserCore(signUpRequest))
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	authResponse := dto.AuthResponse{
		User:        *mappers.FromUserCoreToUserResponse(user.User),
		AccessToken: user.AccessToken,
	}
	c.JSON(http.StatusCreated, authResponse)
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	if err := c.BindJSON(&signInRequest); err != nil {
		utils.HTTPError(errors.ErrBadRequest, c)
		return
	}

	user, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	authResponse := dto.AuthResponse{
		User:        *mappers.FromUserCoreToUserResponse(user.User),
		AccessToken: user.AccessToken,
	}
	c.JSON(http.StatusCreated, authResponse)
}

func (uc *UserHandler) AddUserAddresses(c *gin.Context) {
	var userAddressesRequest dto.UserAddressesRequest
	if err := c.BindJSON(&userAddressesRequest); err != nil {
		utils.HTTPError(errors.ErrBadRequest, c)
		return
	}

	addresses, err := uc.userService.AddUserAddresses(
		userAddressesRequest.UserID,
		mappers.FromAddressRequestSliceToAddressCoreSlice(userAddressesRequest.Addresses))
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	// todo: send custom response
	c.JSON(http.StatusCreated, addresses)
}

func (uc *UserHandler) AddUserPaymentCards(c *gin.Context) {
	var userCardsRequest dto.UserPaymentCardsRequest
	if err := c.BindJSON(&userCardsRequest); err != nil {
		utils.HTTPError(errors.ErrBadRequest, c)
		return
	}

	cards, err := uc.userService.AddUserPaymentCard(
		userCardsRequest.UserID,
		mappers.FromPaymentCardRequestSliceToPaymentCardCoreSlice(userCardsRequest.PaymentCards))
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	// todo: send custom response
	c.JSON(http.StatusCreated, cards)
}

func (uc *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	userResponse := make([]dto.UserResponse, 0)
	for _, k := range users {
		userResponse = append(userResponse, *mappers.FromUserCoreToUserResponse(k))
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserHandler) FindUserByID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		utils.HTTPError(errors.ErrBadRequest, c)
		return
	}

	user, err := uc.userService.FindUserByID(userID)
	if err != nil {
		utils.HTTPError(err, c)
		return
	}

	c.JSON(http.StatusOK, *mappers.FromUserCoreToUserResponse(*user))
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}
