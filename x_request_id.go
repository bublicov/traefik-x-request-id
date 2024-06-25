package traefik_x_request_id

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// RequestID a RequestID plugin.
type RequestID struct {
	next http.Handler
	name string
}

// New created a new RequestID plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RequestID{
		next: next,
		name: name,
	}, nil
}

// ServeHTTP implements the http.Handler interface.
func (a *RequestID) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Generate a unique request ID
	requestID := uuid.New().String()

	// Set the request ID in the request header
	req.Header.Set("X-Request-ID", requestID)

	//// TEST: Set the request ID in the response header
	//rw.Header().Set("X-Request-ID", requestID)

	// Call the next handler
	a.next.ServeHTTP(rw, req)
}
