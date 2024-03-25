package repository

import (
	"database/sql"
	"errors"
	"github.com/Bhavya022/GeoData_App/Backend/internal/models"
)

// GeoDataRepository handles database operations related to geospatial data
type GeoDataRepository struct {
	db *sql.DB
}

// NewGeoDataRepository creates a new instance of GeoDataRepository
func NewGeoDataRepository(db *sql.DB) *GeoDataRepository {
	return &GeoDataRepository{db: db}
}

// SaveGeoData saves geospatial data to the database
func (r *GeoDataRepository) SaveGeoData(geoData *models.GeoData) error {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("INSERT INTO geospatial_data (user_id, file_name, file_type, file_content) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute SQL statement
	_, err = stmt.Exec(geoData.UserID, geoData.FileName, geoData.FileType, geoData.FileContent)
	if err != nil {
		return err
	}

	return nil
}

// GetGeoDataByID retrieves geospatial data from the database by ID
func (r *GeoDataRepository) GetGeoDataByID(id int) (*models.GeoData, error) {
	// Prepare SQL statement
	stmt, err := r.db.Prepare("SELECT id, user_id, file_name, file_type, file_content FROM geospatial_data WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute SQL statement
	row := stmt.QueryRow(id)

	// Parse query result
	var geoData models.GeoData
	err = row.Scan(&geoData.ID, &geoData.UserID, &geoData.FileName, &geoData.FileType, &geoData.FileContent)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("geospatial data not found")
		}
		return nil, err
	}

	return &geoData, nil
}

// Add more methods as needed for CRUD operations on geospatial data
