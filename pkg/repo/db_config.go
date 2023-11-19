package repo

// DatabaseConfig holds configuration for database connections.
type DatabaseConfig struct {
	ConnectionString string `json:"connectionString"`
	DatabaseName     string `json:"databaseName"`
}
