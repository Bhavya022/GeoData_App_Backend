package auth

import (
	"errors"
	"github.com/Bhavya022/GeoData_App/Backend/internal/models"
	"github.com/Bhavya022/GeoData_App/Backend/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

// AuthService handles user authentication
type AuthService struct {
	userRepository *repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

// AuthenticateUser authenticates a user by user ID and password
func (s *AuthService) AuthenticateUser(userID int, password string) (*models.User, error) {
	// Retrieve user by ID
	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Check if user exists
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if password matches
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

// VerifyToken verifies the authenticity of the provided token
func (s *AuthService) VerifyToken(r *http.Request) bool {
	// Extract token from request headers
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// You might need to provide your signing key or fetch it from a secure location
		return []byte("your-secret-key"), nil
	})
	if err != nil || !token.Valid {
		return false
	}

	return true
}
