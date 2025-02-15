package userinterface

import (
	"context"

	"github.com/gofiber/fiber/v2"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, req *usermodel.User) (*usermodel.User, error)
	GetUser(ctx context.Context, phoneNumber, email string) (*usermodel.User, error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (*usermodel.User, error)
	GetUserByID(ctx context.Context, userID int) (*usermodel.User, error)
	UpdateUser(ctx context.Context, userID int, userInput usermodel.UpdateUser) error
}

type UserServiceInterface interface {
	Register(ctx context.Context, req usermodel.User) error
	Login(ctx context.Context, req usermodel.LoginRequest) (*usermodel.LoginResponse, error)

	GetProfile(ctx context.Context, userID int) (*usermodel.User, error)
	UpdateUser(ctx context.Context, userID int, req usermodel.UpdateUser) error
}

type UserHandlerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error

	GetProfile(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}
