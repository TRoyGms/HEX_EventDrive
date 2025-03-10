package database
 
import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)
 
// NewDBConnection creates and returns a new MySQL database connection
func NewDBConnection() (*sql.DB, error) {
    // Get the MySQL DSN (Data Source Name) from the environment
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        return nil, fmt.Errorf("MYSQL_DSN environment variable is not set")
    }

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
