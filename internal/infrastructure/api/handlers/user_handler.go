package handlers

import (
	"github.com/rierarizzo/cafelatte/internal/core/constants"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	errorHandler "github.com/rierarizzo/cafelatte/internal/infrastructure/api/error"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/api/mappers"
	"github.com/rierarizzo/cafelatte/internal/utils"
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
	if errorHandler.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignUp(*mappers.FromSignUpReqToUser(signUpRequest))
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusCreated, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func (uc *UserHandler) SignIn(c *gin.Context) {
	var signInRequest dto.SignInRequest
	err := c.BindJSON(&signInRequest)
	if errorHandler.Error(c, err) {
		return
	}

	authorizedUser, err := uc.userService.SignIn(signInRequest.Email, signInRequest.Password)
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusCreated, mappers.FromAuthorizedUserToAuthorizationRes(*authorizedUser))
}

func getUserIDFromClaims(c *gin.Context) (int, error) {
	userClaims, exists := c.Get(constants.UserClaimsKey)
	if !exists {
		return 0, errors.WrapError(errors.ErrUnauthorizedUser, "claims not present in token")
	}

	return userClaims.(*utils.UserClaims).ID, nil
}

func (uc *UserHandler) AddUserAddresses(c *gin.Context) {
	var addressesRequest []dto.AddressRequest
	err := c.BindJSON(&addressesRequest)
	if errorHandler.Error(c, err) {
		return
	}

	addresses := make([]entities.Address, 0)
	for _, v := range addressesRequest {
		addresses = append(addresses, *mappers.FromAddressReqToAddress(v))
	}

	userID, err := getUserIDFromClaims(c)
	if errorHandler.Error(c, err) {
		return
	}

	addresses, err = uc.userService.AddUserAddresses(userID, addresses)
	if errorHandler.Error(c, err) {
		return
	}

	addressesRes := make([]dto.AddressResponse, 0)
	for _, v := range addresses {
		addressesRes = append(addressesRes, *mappers.FromAddressToAddressRes(v))
	}

	c.JSON(http.StatusCreated, addressesRes)
}

func (uc *UserHandler) AddUserPaymentCards(c *gin.Context) {
	var cardsRequest []dto.PaymentCardRequest
	err := c.BindJSON(&cardsRequest)
	if errorHandler.Error(c, err) {
		return
	}

	cards := make([]entities.PaymentCard, 0)
	for _, v := range cardsRequest {
		cards = append(cards, *mappers.FromPaymentCardReqToPaymentCard(v))
	}

	userID, err := getUserIDFromClaims(c)
	if errorHandler.Error(c, err) {
		return
	}

	cards, err = uc.userService.AddUserPaymentCard(userID, cards)
	if errorHandler.Error(c, err) {
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
	if errorHandler.Error(c, err) {
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
	if errorHandler.Error(c, err) {
		return
	}

	user, err := uc.userService.FindUserByID(userID)
	if errorHandler.Error(c, err) {
		return
	}

	c.JSON(http.StatusOK, *mappers.FromUserToUserRes(*user))
}

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{userService}
}
