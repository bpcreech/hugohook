// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/bpcreech/hugohook/restapi/operations"
    "fmt"
    "bytes"
    "io/ioutil"
    "github.com/google/go-github/github"
    "os"
)

//go:generate swagger generate server --target .. --name  --spec ../hugohook.yaml

func configureFlags(api *operations.HugohookAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HugohookAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HandleEventHandler = operations.HandleEventHandlerFunc(operations.HandleEventHandlerImpl);

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	// Use a custom http.Handler which does HMAC authentication
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    	secret := os.Getenv("HUGOHOOK_SECRET")

        payload, err := github.ValidatePayload(r, []byte(secret))
        if err != nil {
	        fmt.Println("fail to validate request payload")

			rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

			rw.WriteHeader(401)
        }

        r.Body = ioutil.NopCloser(bytes.NewReader(payload))

        handler.ServeHTTP(rw, r)
    })
	// return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
