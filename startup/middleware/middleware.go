package middleware

import (
	"context"
	"net/http"
	"sync"

	"github.com/dev-crusader404/go-test-project/models"
	"github.com/dev-crusader404/go-test-project/startup"
	"github.com/google/uuid"
)

var (
	RequestIDKey = models.RequestID("requestID")
	mutex        sync.Mutex
)

func GenerateRequestID() string {
	mutex.Lock()
	id := uuid.New()
	mutex.Unlock()
	return id.String()
}

func MethodType(next http.HandlerFunc, methodType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if methodType != r.Method {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		next(w, r)
	}
}
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		expectedUserName := startup.GetAll().GetString("BASIC-LOGIN", "")
		expectedPass := startup.GetAll().GetString("BASIC-PASSWORD", "")
		if username != expectedUserName || password != expectedPass {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

// Middleware function for request logging
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = GenerateRequestID()
		}

		// Attach requestID to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, requestID)

		// Call the next handler with the updated request context
		r = r.WithContext(ctx)
		next(w, r)
	}
}
