package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	alamathandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/alamat_handler"
	healthhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/health_handler"
	userhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/user_handler"
	wilayahhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/wilayah_handler"
	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	wilayahinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/wilayah_interface"
	alamatrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/alamat_repository"
	userrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/user_repository"
	wilayahrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/wilayah_repository"
	alamatservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/alamat_service"
	healthservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/health_service"
	userservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/user_service"
	wilayahservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/wilayah_service"
	"gorm.io/gorm"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject(helpers.DB)

	api := app.Group("/api/v1")
	api.Get("/health", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.healthHandler.HealthCheck)

	api.Post("/auth/register", dependency.userHandler.Register)
	api.Post("/auth/login", dependency.userHandler.Login)

	api.Get("/provcity/listprovincies", dependency.wilayahHandler.GetProvinces)
	api.Get("/provcity/listcities/:prov_id", dependency.wilayahHandler.GetCities)
	api.Get("/provcity/detailprovince/:prov_id", dependency.wilayahHandler.GetProvinceDetail)
	api.Get("/provcity/detailcity/:city_id", dependency.wilayahHandler.GetCityDetail)

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

	healthHandler  healthinterface.HealthHandlerInterface
	wilayahHandler wilayahinterface.WilayahHandlerInterface
	userHandler    userinterface.UserHandlerInterface
	alamatHandler  alamatinterface.AlamatHandlerInterface
}

func dependencyInject(db *gorm.DB) Dependency {
	healthSvc := &healthservice.HealthService{}
	healthHandler := &healthhandler.HealthHandler{Service: healthSvc}

	wilayahRepo := &wilayahrepository.WilayahRepository{}
	wilayahSvc := &wilayahservice.WilayahService{WilayahRepository: wilayahRepo}
	wilayahHandler := &wilayahhandler.WilayahHandler{WilayahService: wilayahSvc}

	userRepo := &userrepository.UserRepository{DB: db}
	userSvc := &userservice.UserService{UserRepository: userRepo, WilayahRepository: wilayahRepo}
	userHandler := &userhandler.UserHandler{UserService: userSvc}

	alamatRepo := &alamatrepository.AlamatRepository{DB: db}
	alamatSvc := &alamatservice.AlamatService{AlamatRepository: alamatRepo}
	alamatHandler := &alamathandler.AlamatHandler{AlamatService: alamatSvc}

	return Dependency{
		userRepository: userRepo,
		healthHandler:  healthHandler,
		wilayahHandler: wilayahHandler,
		userHandler:    userHandler,
		alamatHandler:  alamatHandler,
	}
}
