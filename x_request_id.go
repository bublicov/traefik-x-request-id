package traefik_x_request_id

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// Config the plugin configuration.
type Config struct {
	RequestIDType     string `json:"requestIDType"`     // "uuid", "ulid", or "snowflake"
	RequestIDOverride bool   `json:"requestIDOverride"` // Whether to override existing X-Request-ID
	NodeID            int64  `json:"nodeID"`            // Node ID for Snowflake
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		RequestIDType:     "uuid", // default to UUID
		RequestIDOverride: false,  // default to not override existing X-Request-ID
		NodeID:            1,      // default Node ID for Snowflake
	}
}

// RequestID a RequestID plugin.
type RequestID struct {
	next   http.Handler
	name   string
	config *Config
	node   *snowflake.Node
}

// New created a new RequestID plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	var node *snowflake.Node
	var err error

	if config.RequestIDType == "snowflake" {
		node, err = snowflake.NewNode(config.NodeID) // Use the configured Node ID
		if err != nil {
			return nil, err
		}
	}

	return &RequestID{
		next:   next,
		name:   name,
		config: config,
		node:   node,
	}, nil
}

// ServeHTTP implements the http.Handler interface.
func (a *RequestID) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var requestID string

	// Check if X-Request-ID already exists and should not be overridden
	if req.Header.Get("X-Request-ID") == "" || a.config.RequestIDOverride {
		switch a.config.RequestIDType {
		case "ulid":
			entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
			id, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
			if err != nil {
				// Handle error, possibly log it and fall back to UUID
				requestID = generateUUID()
			} else {
				requestID = id.String()
			}
		case "snowflake":
			if a.node != nil {
				requestID = a.node.Generate().String()
			} else {
				// Handle error, possibly log it and fall back to UUID
				requestID = generateUUID()
			}
		default:
			requestID = generateUUID()
		}

		// Set the request ID in the request header
		req.Header.Set("X-Request-ID", requestID)

		//// TEST: Set the request ID in the response header
		//rw.Header().Set("X-Request-ID", requestID)
	}

	// Call the next handler
	a.next.ServeHTTP(rw, req)
}

/* Helpers
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
// generateUUID generates a new UUID.
func generateUUID() string {
	return uuid.New().String()
}
