package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/external"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	alamathandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/alamat_handler"
	healthhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/health_handler"
	userhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/user_handler"
	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	alamatrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/alamat_repository"
	userrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/user_repository"
	alamatservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/alamat_service"
	healthservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/health_service"
	userservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/user_service"
	"gorm.io/gorm"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject(helpers.DB)

	api := app.Group("/api/v1")
	api.Get("/health", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.healthHandler.HealthCheck)

	api.Post("/auth/register", dependency.userHandler.Register)
	api.Post("/auth/login", dependency.userHandler.Login)

	api.Get("/user", dependency.AuthMiddleware, dependency.userHandler.GetProfile)
	api.Put("/user", dependency.AuthMiddleware, dependency.userHandler.UpdateUser)
	api.Post("/user/alamat", dependency.AuthMiddleware, dependency.alamatHandler.CreateAlamat)
	api.Get("/user/alamat", dependency.AuthMiddleware, dependency.alamatHandler.GetAlamat)
	api.Get("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.GetAlamatByID)
	api.Put("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.UpdateAlamat)
	api.Delete("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.DeleteAlamat)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	userRepository userinterface.UserRepositoryInterface

	healthHandler healthinterface.HealthHandlerInterface
	userHandler   userinterface.UserHandlerInterface
	alamatHandler alamatinterface.AlamatHandlerInterface
}

func dependencyInject(db *gorm.DB) Dependency {
	external := &external.External{}

	healthSvc := &healthservice.HealthService{}
	healthHandler := &healthhandler.HealthHandler{Service: healthSvc}

	userRepo := &userrepository.UserRepository{DB: db}
	userSvc := &userservice.UserService{UserRepository: userRepo, External: external}
	userHandler := &userhandler.UserHandler{UserService: userSvc}

	alamatRepo := &alamatrepository.AlamatRepository{DB: db}
	alamatSvc := &alamatservice.AlamatService{AlamatRepository: alamatRepo}
	alamatHandler := &alamathandler.AlamatHandler{AlamatService: alamatSvc}

	return Dependency{
		userRepository: userRepo,
		healthHandler:  healthHandler,
		userHandler:    userHandler,
		alamatHandler:  alamatHandler,
	}
}
