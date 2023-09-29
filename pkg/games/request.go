package games

// Define a struct called CreateRequest
type CreateRequest struct {
	// Name field of type string with json tag "name" and binding tag "required"
	Name string `json:"name" binding:"required"`
	// Genre field of type string with json tag "genre" and binding tag "required"
	Genre string `json:"genre" binding:"required"`
	// ReleaseDate field of type string with json tag "release_date" and binding tag "required,datetime=2006-01-02"
	// The datetime tag specifies the format of the date string as "2006-01-02"
	ReleaseDate string `json:"release_date" binding:"required,datetime=2006-01-02"`
}
