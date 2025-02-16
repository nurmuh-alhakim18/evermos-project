package userrepository

import (
	"context"

	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(ctx context.Context, req *usermodel.User) (*usermodel.User, error) {
	err := r.DB.WithContext(ctx).Create(req).Error
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *UserRepository) GetUser(ctx context.Context, phoneNumber, email string) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).Where("no_telp = ? AND email = ?", phoneNumber, email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phoneNumber string) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).Where("no_telp = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID int) (*usermodel.User, error) {
	var user usermodel.User
	err := r.DB.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userID int, userInput usermodel.UpdateUser) error {
	var user usermodel.User
	err := r.DB.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return err
	}

	return r.DB.WithContext(ctx).Model(&user).Updates(userInput).Error
}
