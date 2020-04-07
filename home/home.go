package home

import (
	"encoding/json"
	"log"
	"net/http"
)

const genericResponse = "Welcome to modb-server!"

// HomepageResponse is a type defining the structure of the returned JSON.
// @TODO This needs to go and reside together with all Swagger-defined models somewhere else.
type HomepageResponse struct {
	Message string `json:"message"`
}

// Handlers is a type enabling dependency injection.
type Handlers struct {
	logger *log.Logger
}

// Home is the handler for the home "/" route of this server.
// It's just returning a generic greeting message.
func (handlers *Handlers) Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	response := HomepageResponse{
		Message: genericResponse,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		handlers.logger.Fatalf("JSON could not be marshalled: %v", err)
	} else {
		defer handlers.logger.Printf("Client %v requested Home.", request.RemoteAddr)
		writer.Write([]byte(responseJSON))
	}
}

// Logger (doc tbd.)
func (handlers *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		next(writer, request)
	}
}

// SetupRoutes sets up the routes...
func (handlers *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.Logger(handlers.Home))
}

// NewHandlers returns a pointer to a new created handler for this class.
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
