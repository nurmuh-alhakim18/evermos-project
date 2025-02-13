package userrepository

import (
	"context"
	"errors"

	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(ctx context.Context, user *usermodel.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUser(ctx context.Context, phoneNumber, email string) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).Where("no_telp = ? AND email = ?", phoneNumber, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phoneNumber string) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.Where("no_telp = ?", phoneNumber).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID int) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
