package repository

import (
	"database/sql"
	"errors"

	"github.com/Bhavya022/GeoData_App/Backend/internal/models"
)

// UserRepository handles database operations related to users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("INSERT INTO users (username, email) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(user.Username, user.Email)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByID retrieves a user from the database by ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("SELECT id, username, email FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute SQL statement
	row := stmt.QueryRow(id)

	// Parse query result
	var user models.User
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *models.User) error {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("UPDATE users SET username = $1, email = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user from the database by ID
func (r *UserRepository) DeleteUser(id int) error {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
