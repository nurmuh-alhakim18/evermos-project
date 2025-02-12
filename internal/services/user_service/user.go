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
		return fmt.Errorf("failed to check if user exists: %w", err)
	}

	if user != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := req
	newUser.KataSandi = string(hashedPassword)

	err = s.UserRepository.CreateUser(ctx, &newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, req usermodel.LoginRequest) (*usermodel.LoginResponse, error) {
	user, err := s.UserRepository.GetUserByPhone(ctx, req.NoTelp)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not exists: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(req.KataSandi))
	if err != nil {
		return nil, fmt.Errorf("failed to compare password: %w", err)
	}

	token, err := helpers.GenerateJWT(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	provinsi, err := s.External.GetProvince(user.IdProvinsi)
	if err != nil {
		return nil, fmt.Errorf("failed to get province: %w", err)
	}

	kota, err := s.External.GetCity(user.IdProvinsi, user.IdKota)
	if err != nil {
		return nil, fmt.Errorf("failed to get city: %w", err)
	}

	return &usermodel.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		Provinsi:     provinsi,
		Kota:         kota,
		Token:        token,
	}, nil
}
