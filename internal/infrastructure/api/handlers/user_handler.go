package handlers

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/dto"
)

type UserHandler struct {
	userService ports.IUserService
}

func (uc *UserHandler) SignUp(c *gin.Context) {
	var signUpRequest dto.SignUpRequest
	err := c.BindJSON(&signUpRequest)
	if error.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignUp(*mappers.FromSignUpReqToUser(signUpRequest))
	if error.Error(c, err) {
		return
	}

	c.JSON(http.StatusCreated, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if error.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if error.Error(c, err) {
		return
	}

	c.JSON(http.StatusCreated, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (uc *UserHandler) AddUserAddresses(c *gin.Context) {
	var userAddressesRequest dto.UserAddressesRequest
	err := c.BindJSON(&userAddressesRequest)
	if error.Error(c, err) {
		return
	}

	addresses := make([]entities.Address, 0)
	for _, v := range userAddressesRequest.Addresses {
		addresses = append(addresses, *mappers.FromAddressReqToAddress(v))
	}

	addresses, err = uc.userService.AddUserAddresses(userAddressesRequest.UserID, addresses)
	if error.Error(c, err) {
		return
	}

	addressesRes := make([]dto.AddressResponse, 0)
	for _, v := range addresses {
		addressesRes = append(addressesRes, *mappers.FromAddressToAddressRes(v))
	}

	c.JSON(http.StatusCreated, addressesRes)
}

func (uc *UserHandler) AddUserPaymentCards(c *gin.Context) {
	var userCardsRequest dto.UserPaymentCardsRequest
	err := c.BindJSON(&userCardsRequest)
	if error.Error(c, err) {
		return
	}

	cards := make([]entities.PaymentCard, 0)
	for _, v := range userCardsRequest.PaymentCards {
		cards = append(cards, *mappers.FromPaymentCardReqToPaymentCard(v))
	}

	cards, err = uc.userService.AddUserPaymentCard(userCardsRequest.UserID, cards)
	if error.Error(c, err) {
		return
	}

	cardsRes := make([]dto.PaymentCardResponse, 0)
	for _, v := range cards {
		cardsRes = append(cardsRes, *mappers.FromPaymentCardToPaymentCardRes(v))
	}

	c.JSON(http.StatusCreated, cardsRes)
}

func (uc *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if error.Error(c, err) {
		return
	}

	userResponse := make([]dto.UserResponse, 0)
	for _, k := range users {
		userResponse = append(userResponse, *mappers.FromUserToUserRes(k))
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserHandler) FindUserByID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if error.Error(c, err) {
		return
	}

	user, err := uc.userService.FindUserByID(userID)
	if error.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, *mappers.FromUserToUserRes(*user))
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}
