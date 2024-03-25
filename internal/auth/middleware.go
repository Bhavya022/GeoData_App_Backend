package auth

import (
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware to check if the user is authenticated
type AuthMiddleware struct {
	AuthService *AuthService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(authService *AuthService) *AuthMiddleware {
	return &AuthMiddleware{AuthService: authService}
}

// MiddlewareFunc is a function signature for middleware
type MiddlewareFunc func(http.Handler) http.Handler

// Middleware is the middleware implementation
// Middleware is the middleware implementation
func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify token and check if user is authenticated
		if !m.AuthService.VerifyToken(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If user is authenticated, call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Extract authentication token from request headers, cookies, or query parameters
// 		token := extractToken(r)

// 		// Verify token and check if user is authenticated
// 		if !m.AuthService.VerifyToken(token) {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// If user is authenticated, call the next handler in the chain
// 		next.ServeHTTP(w, r)
// 	})
// }

// extractToken extracts the authentication token from the request
// func extractToken(r *http.Request) string {
// 	// Extract token from request headers
// 	token := r.Header.Get("Authorization")

// 	// If token is empty, try extracting from cookies
// 	if token == "" {
// 		cookie, err := r.Cookie("auth_token")
// 		if err == nil {
// 			token = cookie.Value
// 		}
// 	}

// 	// If token is still empty, try extracting from query parameters
// 	if token == "" {
// 		token = r.URL.Query().Get("token")
// 	}

// 	return token
// }

// extractToken extracts the authentication token from the request
func extractToken(r *http.Request) string {
	// Extract token from request headers
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// Check if the Authorization header contains a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}

	// If token is empty, try extracting from cookies
	token := ""
	cookie, err := r.Cookie("auth_token")
	if err == nil {
		token = cookie.Value
	}

	// If token is still empty, try extracting from query parameters
	if token == "" {
		token = r.URL.Query().Get("token")
	}

	return token
}
