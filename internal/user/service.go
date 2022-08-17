package user

import (
	"context"
	"fmt"
	"golang-rest-api/pkg/logging"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *Service) GetAllUsers(ctx context.Context) (u []*User, err error) {
	u, err = s.storage.FindAll(ctx)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) GetOneUser(ctx context.Context, id string) (u *User, err error) {
	u, err = s.storage.FindOne(ctx, id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) CreateUser(ctx context.Context, dto CreateUserDTO) (u *User, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)

	if err != nil {
		return u, fmt.Errorf("failed to hash password due to error: %v", err)
	}

	user := User{
		ID:           "",
		Username:     dto.Username,
		PasswordHash: string(passwordHash),
		Email:        dto.Email,
	}

	u, err = s.storage.Create(ctx, &user)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) UpdateUser(ctx context.Context, id string, dto UpdateUserDTO) (u *User, err error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 8)
	if err != nil {
		return u, fmt.Errorf("failed to hash password due to error: %v", err)
	}

	user := User{
		ID:           id,
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: string(passwordHash),
	}

	u, err = s.storage.Update(ctx, &user)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) (err error) {
	err = s.storage.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewService(storage Storage, logger *logging.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}
