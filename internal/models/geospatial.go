package models

// GeoData represents geospatial data uploaded by users
type GeoData struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	FileContent []byte `json:"fileContent"`
	// Add more fields as needed
}

// Shape represents a custom shape drawn by users on the map
type Shape struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userID"`
	Geometry []byte `json:"geometry"`
	// Add more fields as needed
}
