// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kaatinga/plantbook/restapi/operations"
	"github.com/kaatinga/plantbook/restapi/operations/plant"
	"github.com/kaatinga/plantbook/restapi/operations/user"
)

//go:generate swagger generate server --target ../../plantbook --name Plantbook --spec ../api/swagger.yaml --principal interface{}

func configureFlags(api *operations.PlantbookAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PlantbookAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer
	api.UrlformConsumer = runtime.DiscardConsumer
	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.XMLProducer = runtime.XMLProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// plant.UpdatePlantWithFormMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// plant.UploadFileMaxParseMemory = 32 << 20

	if api.PlantAddPlantHandler == nil {
		api.PlantAddPlantHandler = plant.AddPlantHandlerFunc(func(params plant.AddPlantParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.AddPlant has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.PlantDeletePlantHandler == nil {
		api.PlantDeletePlantHandler = plant.DeletePlantHandlerFunc(func(params plant.DeletePlantParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.DeletePlant has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.PlantGetPlantByIDHandler == nil {
		api.PlantGetPlantByIDHandler = plant.GetPlantByIDHandlerFunc(func(params plant.GetPlantByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.GetPlantByID has not yet been implemented")
		})
	}
	if api.UserGetUserByNameHandler == nil {
		api.UserGetUserByNameHandler = user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
		})
	}
	if api.UserLoginUserHandler == nil {
		api.UserLoginUserHandler = user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		})
	}
	if api.UserLogoutUserHandler == nil {
		api.UserLogoutUserHandler = user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		})
	}
	if api.PlantUpdatePlantHandler == nil {
		api.PlantUpdatePlantHandler = plant.UpdatePlantHandlerFunc(func(params plant.UpdatePlantParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.UpdatePlant has not yet been implemented")
		})
	}
	if api.PlantUpdatePlantWithFormHandler == nil {
		api.PlantUpdatePlantWithFormHandler = plant.UpdatePlantWithFormHandlerFunc(func(params plant.UpdatePlantWithFormParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.UpdatePlantWithForm has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}
	if api.PlantUploadFileHandler == nil {
		api.PlantUploadFileHandler = plant.UploadFileHandlerFunc(func(params plant.UploadFileParams) middleware.Responder {
			return middleware.NotImplemented("operation plant.UploadFile has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
