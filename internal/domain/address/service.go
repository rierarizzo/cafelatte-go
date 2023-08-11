package address

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type Service struct {
	addressRepo IAddressRepository
}

func (s Service) GetAddressesByUserID(userID int) ([]Address, *domain.AppError) {
	addresses, appErr := s.addressRepo.SelectAddressesByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func (s Service) AddUserAddresses(userID int,
	addresses []Address) ([]Address, *domain.AppError) {
	for _, v := range addresses {
		if appErr := ValidateAddress(&v); appErr != nil {
			return nil, appErr
		}
	}

	addresses, appErr := s.addressRepo.InsertUserAddresses(userID, addresses)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func NewAddressService(addressRepo IAddressRepository) *Service {
	return &Service{addressRepo}
}
