package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bhavya022/GeoData_App/Backend/internal/auth"
	"github.com/Bhavya022/GeoData_App/Backend/internal/models"
	"github.com/Bhavya022/GeoData_App/Backend/internal/repository"
	"github.com/go-chi/chi"
)

// type LoginRequest struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// LoginHandler handles user login requests
// func LoginHandler(authService *auth.AuthService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var loginRequest LoginRequest
// 		err := json.NewDecoder(r.Body).Decode(&loginRequest)
// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}

// 		// Authenticate user
// 		user, err := authService.AuthenticateUser(loginRequest.Username, loginRequest.Password)
// 		if err != nil {
// 			http.Error(w, "Authentication failed", http.StatusUnauthorized)
// 			return
// 		}

// 		// Respond with user data (you may want to exclude sensitive fields like password)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(user)
// 	}
// }

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles user login requests
func LoginHandler(authService *auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginRequest LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Convert the username to an integer (assuming it's an ID)
		userID, err := strconv.Atoi(loginRequest.Username)
		if err != nil {
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		}

		// Authenticate user using user ID and password
		user, err := authService.AuthenticateUser(userID, loginRequest.Password)
		if err != nil {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}

		// Respond with user data (you may want to exclude sensitive fields like password)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo *repository.UserRepository
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// GetUserByIDHandler handles requests to retrieve a user by ID
func (h *UserHandler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request parameters
	userID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get user by ID from the repository
	user, err := h.userRepo.GetUserByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// Check if user exists
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUserHandler handles requests to create a new user
func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userRepo.CreateUser(&user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateUserHandler handles requests to update an existing user
func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request parameters
	userID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the request body into a User struct
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update the user in the repository
	user.ID = id
	err = h.userRepo.UpdateUser(&user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUserHandler handles requests to delete a user by ID
func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request parameters
	userID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete the user from the repository
	err = h.userRepo.DeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
