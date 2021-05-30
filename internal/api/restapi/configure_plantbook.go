// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaatinga/plantbook/pkg/logging"
	"github.com/kaatinga/plantbook/pkg/token"

	hhandlers "github.com/kaatinga/plantbook/internal/api/handlers/health"
	phandlers "github.com/kaatinga/plantbook/internal/api/handlers/plants"
	uhandlers "github.com/kaatinga/plantbook/internal/api/handlers/users"
	apimiddleware "github.com/kaatinga/plantbook/internal/api/middleware"
	"github.com/kaatinga/plantbook/internal/api/repo"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/plant"
	"github.com/kaatinga/plantbook/internal/api/restapi/operations/user"
	"github.com/kaatinga/plantbook/internal/config"
	"github.com/kaatinga/plantbook/internal/metric"
)

//go:generate swagger generate server --target ../../api --name Plantbook --spec ../../../api/swagger/swagger.yaml --principal interface{}

const (
	version          string        = "0.0.2"
	tokenExpireDelay time.Duration = 7 * 24 * 60 * time.Minute
)

var (
	build   string = "_build_ldflags"
	githash string = "_githash_ldflags"
	buildAt string = "_git_build_at_ldflags"
	cfg     config.Config
)

func configureFlags(api *operations.PlantbookAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PlantbookAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// read config
	err := config.Read(&config.Defaults, &cfg)
	if err != nil {
		log.Fatalf("read config fatal error, %s", err)
	}
	// set logger
	logger := logging.NewLogger(cfg.LOG.Debug, cfg.LOG.Format)
	logger = logger.With("version", version)
	logger = logger.With("build", build)
	logger = logger.With("githash", githash)
	logger = logger.With("build_at", buildAt)

	ctx := logging.WithLogger(context.Background(), logger)
	// make repo
	repo, err := repo.NewPG(ctx, cfg.DB.URL, cfg.LOG.Debug)
	if err != nil {
		logger.Fatalf("connect to storage error, %s", err)
	}
	tm, err := token.NewJwtToken(cfg.TokenSecret)
	if err != nil {
		logger.Fatalf("make token manager error, %s", err)
	}
	// make handlers
	// health
	api.HealthHealthAliveHandler = hhandlers.NewHealthAliveHandler()
	api.HealthHealthReadyHandler = hhandlers.NewHealthReadyHandler(repo)

	buildAtTime, err := time.Parse(time.RFC3339, buildAt)
	if err != nil {
		logger.Warnf("incorrect buildAt %s, set to current time", err)
		buildAtTime = time.Now()
	}
	api.HealthAPIVersionHandler = hhandlers.NewAPIVersionHandler(version, githash, buildAtTime)
	// metrics

	// users
	api.UserCreateUserHandler = uhandlers.NewCreateUserHandler(repo, tm)
	api.UserLoginUserHandler = uhandlers.NewLoginUserHandler(repo, tm, tokenExpireDelay)
	api.UserLogoutUserHandler = uhandlers.NewLogoutUserHandler(tokenExpireDelay)

	// plants TODO: fill me
	api.PlantGetPlantsHandler = phandlers.NewGetPlantsHandler(repo)
	//

	// generated code...
	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	// api.MultipartformConsumer = runtime.DiscardConsumer
	// api.UrlformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// plant.UpdatePlantWithFormMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// plant.UploadFileMaxParseMemory = 32 << 20

	// if api.PlantAddPlantHandler == nil {
	// 	api.PlantAddPlantHandler = plant.AddPlantHandlerFunc(func(params plant.AddPlantParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation plant.AddPlant has not yet been implemented")
	// 	})
	// }

	// if api.PlantDeletePlantHandler == nil {
	// 	api.PlantDeletePlantHandler = plant.DeletePlantHandlerFunc(func(params plant.DeletePlantParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation plant.DeletePlant has not yet been implemented")
	// 	})
	// }
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.PlantGetPlantByIDHandler == nil {
		api.PlantGetPlantByIDHandler = plant.GetPlantByIDHandlerFunc(
			func(params plant.GetPlantByIDParams) middleware.Responder {
				return middleware.NotImplemented("operation plant.GetPlantByID has not yet been implemented")
			})
	}
	if api.UserGetUserByNameHandler == nil {
		api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(
			func(params user.GetUserByNameParams) middleware.Responder {
				return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
			})
	}
	// if api.PlantUpdatePlantHandler == nil {
	// 	api.PlantUpdatePlantHandler = plant.UpdatePlantHandlerFunc(func(params plant.UpdatePlantParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation plant.UpdatePlant has not yet been implemented")
	// 	})
	// }
	// if api.PlantUpdatePlantWithFormHandler == nil {
	// 	api.PlantUpdatePlantWithFormHandler = plant.UpdatePlantWithFormHandlerFunc(
	// 		func(params plant.UpdatePlantWithFormParams) middleware.Responder {
	// 			return middleware.NotImplemented("operation plant.UpdatePlantWithForm has not yet been implemented")
	// 		})
	// }
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}
	// if api.PlantUploadFileHandler == nil {
	// 	api.PlantUploadFileHandler = plant.UploadFileHandlerFunc(func(params plant.UploadFileParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation plant.UploadFile has not yet been implemented")
	// 	})
	// }

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	// metrics
	metricSRV := metric.NewServer(cfg.HTTPD.Host + ":" + cfg.HTTPD.MetricPort)
	go metricSRV.Run(ctx)

	return setupGlobalMiddleware(ctx, api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
// nolint
func configureServer(s *http.Server, scheme, addr string) {
	s.Addr = cfg.HTTPD.Host + ":" + cfg.HTTPD.Port
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything,
// this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	chainGlobalMiddleware := apimiddleware.RequestID(ctx, handler)
	//
	return apimiddleware.SetupHandler(chainGlobalMiddleware)
}
