package userservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	externalinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/external_interface"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository userinterface.UserRepositoryInterface
	External       externalinterface.ExternalInterface
}

func (s *UserService) Register(ctx context.Context, req usermodel.User) error {
	user, err := s.UserRepository.GetUser(ctx, req.NoTelp, req.Email)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %v", err)
	}

	if user != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	newUser := req
	newUser.KataSandi = string(hashedPassword)

	date, err := helpers.ParseBirthDate(newUser.TanggalLahir)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	newUser.TanggalLahir = date.String()

	err = s.UserRepository.CreateUser(ctx, &newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, req usermodel.LoginRequest) (*usermodel.LoginResponse, error) {
	user, err := s.UserRepository.GetUserByPhone(ctx, req.NoTelp)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not exists: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(req.KataSandi))
	if err != nil {
		return nil, fmt.Errorf("failed to compare password: %v", err)
	}

	date, err := helpers.BirthDateToIndoFormat(user.TanggalLahir)
	if err != nil {
		return nil, fmt.Errorf("failed to convert birth date: %v", err)
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	provinsi, err := s.External.GetProvince(user.IdProvinsi)
	if err != nil {
		return nil, fmt.Errorf("failed to get province: %v", err)
	}

	kota, err := s.External.GetCity(user.IdProvinsi, user.IdKota)
	if err != nil {
		return nil, fmt.Errorf("failed to get city: %v", err)
	}

	return &usermodel.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: date,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		Provinsi:     provinsi,
		Kota:         kota,
		Token:        token,
	}, nil
}

func (s *UserService) GetProfile(ctx context.Context, userID int) (*usermodel.User, error) {
	user, err := s.UserRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not exists: %v", err)
	}

	date, err := helpers.BirthDateToIndoFormat(user.TanggalLahir)
	if err != nil {
		return nil, fmt.Errorf("failed to convert birth date: %v", err)
	}

	user.TanggalLahir = date

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userID int, req usermodel.UpdateUser) (*usermodel.UpdateUser, error) {
	date, err := helpers.ParseBirthDate(req.TanggalLahir)
	if err != nil {
		return nil, fmt.Errorf("failed to convert birth date: %v", err)
	}

	userUpdate := req
	userUpdate.TanggalLahir = date.String()

	user, err := s.UserRepository.UpdateUser(ctx, userID, userUpdate)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not exists: %v", err)
	}

	return user, nil
}
