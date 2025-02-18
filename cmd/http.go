package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	alamathandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/alamat_handler"
	healthhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/health_handler"
	kategorihandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/kategori_handler"
	produkhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/produk_handler"
	tokohandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/toko_handler"
	trxhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/trx_handler"
	userhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/user_handler"
	wilayahhandler "github.com/nurmuh-alhakim18/evermos-project/internal/handlers/wilayah_handler"
	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	produkinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/produk_interface"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	trxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/trx_interface"
	userinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/user_interface"
	wilayahinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/wilayah_interface"
	alamatrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/alamat_repository"
	fotoprodukrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/foto_produk_repository"
	kategorirepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/kategori_repository"
	produkrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/produk_repository"
	tokorepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/toko_repository"
	trxrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/trx_repository"
	userrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/user_repository"
	wilayahrepository "github.com/nurmuh-alhakim18/evermos-project/internal/repositories/wilayah_repository"
	alamatservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/alamat_service"
	fotoprodukservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/foto_produk_service"
	healthservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/health_service"
	kategoriservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/kategori_service"
	produkservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/produk_service"
	tokoservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/toko_service"
	trxservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/trx_service"
	userservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/user_service"
	wilayahservice "github.com/nurmuh-alhakim18/evermos-project/internal/services/wilayah_service"
	"gorm.io/gorm"
)

func ServeHTTP() {
	app := fiber.New()
	dependency := dependencyInject(helpers.DB)

	api := app.Group("/api/v1")
	api.Get("/health", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.healthHandler.HealthCheck)

	api.Post("/category", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.kategoriHandler.CreateKategori)
	api.Get("/category", dependency.kategoriHandler.GetKategoris)
	api.Get("/category/:id", dependency.AuthMiddleware, dependency.kategoriHandler.GetKategoriByID)
	api.Put("/category/:id", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.kategoriHandler.UpdateKategori)
	api.Delete("/category/:id", dependency.AuthMiddleware, dependency.AdminMiddleware, dependency.kategoriHandler.DeleteKategori)

	api.Post("/auth/register", dependency.userHandler.Register)
	api.Post("/auth/login", dependency.userHandler.Login)

	api.Get("/provcity/listprovincies", dependency.wilayahHandler.GetProvinces)
	api.Get("/provcity/listcities/:prov_id", dependency.wilayahHandler.GetCities)
	api.Get("/provcity/detailprovince/:prov_id", dependency.wilayahHandler.GetProvinceDetail)
	api.Get("/provcity/detailcity/:city_id", dependency.wilayahHandler.GetCityDetail)

	api.Get("/user", dependency.AuthMiddleware, dependency.userHandler.GetProfile)
	api.Put("/user", dependency.AuthMiddleware, dependency.userHandler.UpdateUser)
	api.Post("/user/alamat", dependency.AuthMiddleware, dependency.alamatHandler.CreateAlamat)
	api.Get("/user/alamat", dependency.AuthMiddleware, dependency.alamatHandler.GetAlamats)
	api.Get("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.GetAlamatByID)
	api.Put("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.UpdateAlamat)
	api.Delete("/user/alamat/:id", dependency.AuthMiddleware, dependency.alamatHandler.DeleteAlamat)

	api.Get("/toko", dependency.AuthMiddleware, dependency.tokoHandler.GetTokos)
	api.Get("/toko/my", dependency.AuthMiddleware, dependency.tokoHandler.GetTokoByUserID)
	api.Get("/toko/:id_toko", dependency.AuthMiddleware, dependency.tokoHandler.GetTokoByID)
	api.Put("/toko/:id_toko", dependency.AuthMiddleware, dependency.tokoHandler.UpdateToko)

	api.Post("/product", dependency.AuthMiddleware, dependency.produkHandler.CreateProduk)
	api.Get("/product", dependency.produkHandler.GetProduks)
	api.Get("/product/:id", dependency.produkHandler.GetProdukByID)
	api.Put("/product/:id", dependency.AuthMiddleware, dependency.produkHandler.UpdateProduk)
	api.Delete("/product/:id", dependency.AuthMiddleware, dependency.produkHandler.DeleteProduk)

	api.Post("/trx", dependency.AuthMiddleware, dependency.trxHandler.CreateTrx)
	api.Get("/trx", dependency.AuthMiddleware, dependency.trxHandler.GetTrxByUserID)
	api.Get("/trx/:id", dependency.AuthMiddleware, dependency.trxHandler.GetTrxByID)

	port := helpers.GetEnv("PORT", "8080")
	app.Listen(":" + port)
}

type Dependency struct {
	userRepository userinterface.UserRepositoryInterface

	healthHandler   healthinterface.HealthHandlerInterface
	wilayahHandler  wilayahinterface.WilayahHandlerInterface
	tokoHandler     tokointerface.TokoHandlerInterface
	userHandler     userinterface.UserHandlerInterface
	alamatHandler   alamatinterface.AlamatHandlerInterface
	kategoriHandler kategoriinterface.KategoriHandlerInterface
	produkHandler   produkinterface.ProdukHandlerInterface
	trxHandler      trxinterface.TrxHandlerInterface
}

func dependencyInject(db *gorm.DB) Dependency {
	healthSvc := &healthservice.HealthService{}
	healthHandler := &healthhandler.HealthHandler{Service: healthSvc}

	wilayahRepo := &wilayahrepository.WilayahRepository{}
	wilayahSvc := &wilayahservice.WilayahService{WilayahRepository: wilayahRepo}
	wilayahHandler := &wilayahhandler.WilayahHandler{WilayahService: wilayahSvc}

	tokoRepo := &tokorepository.TokoRepository{DB: db}
	tokoSvc := &tokoservice.TokoService{TokoRepository: tokoRepo}
	tokoHandler := &tokohandler.TokoHandler{TokoService: tokoSvc}

	userRepo := &userrepository.UserRepository{DB: db}
	userSvc := &userservice.UserService{UserRepository: userRepo, WilayahService: wilayahSvc, TokoService: tokoSvc}
	userHandler := &userhandler.UserHandler{UserService: userSvc}

	alamatRepo := &alamatrepository.AlamatRepository{DB: db}
	alamatSvc := &alamatservice.AlamatService{AlamatRepository: alamatRepo}
	alamatHandler := &alamathandler.AlamatHandler{AlamatService: alamatSvc}

	kategoriRepo := &kategorirepository.KategoriRepository{DB: db}
	kategoriSvc := &kategoriservice.KategoriService{KategoriRepository: kategoriRepo}
	kategoriHandler := &kategorihandler.KategoriHandler{KategoriService: kategoriSvc}

	fotoProdukRepo := &fotoprodukrepository.FotoProdukRepository{DB: db}
	fotoProdukSvc := &fotoprodukservice.FotoProdukService{FotoProdukRepository: fotoProdukRepo}

	produkRepo := &produkrepository.ProdukRepository{DB: db}
	produkSvc := &produkservice.ProdukService{
		ProdukRepository:  produkRepo,
		TokoService:       tokoSvc,
		KategoriService:   kategoriSvc,
		FotoProdukService: fotoProdukSvc,
	}
	produkHandler := &produkhandler.ProdukHandler{
		ProdukService: produkSvc,
		TokoService:   tokoSvc,
	}

	trxRepo := &trxrepository.TrxRepository{DB: db}
	trxSvc := &trxservice.TrxService{TrxRepository: trxRepo}
	trxHandler := &trxhandler.TrxHandler{TrxService: trxSvc}

	return Dependency{
		userRepository:  userRepo,
		healthHandler:   healthHandler,
		wilayahHandler:  wilayahHandler,
		userHandler:     userHandler,
		alamatHandler:   alamatHandler,
		tokoHandler:     tokoHandler,
		kategoriHandler: kategoriHandler,
		produkHandler:   produkHandler,
		trxHandler:      trxHandler,
	}
}
