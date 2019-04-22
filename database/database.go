package database

import (
	"database/sql"
	"fmt"
	"github.com/metaleaf-io/gator/conf"
	"github.com/metaleaf-io/log"
	"io/ioutil"
	"path"

	_ "github.com/lib/pq"
)

// Connect establishes a connection with the database server.
func Connect(appConfig *conf.AppConfig) *sql.DB {
	log.Info("Connecting to database",
		log.String("driver", appConfig.Database.Driver),
		log.String("host", appConfig.Database.Host),
		log.Int16("port", appConfig.Database.Port))

	// This is Go's "connection string"
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.Username,
		appConfig.Database.Password,
		appConfig.Database.Database)

	// Connect
	var err error
	var db  *sql.DB
	if db, err = sql.Open(appConfig.Database.Driver, dbInfo); err != nil {
		panic(err)
	}

	// Verify the connection.
	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

// Migrate verifies the version of the database currently running, then
// applies any newer migration scripts.
//
// db is an active database connection
// dirPath is the path of a directory that contains migration files. The
//      convention is to name the files with an incremental number so that
//      they are loaded from disk in order.
func Migrate(db *sql.DB, dirPath string) {
	if db == nil {
		panic("Nil database connection")
	}

	if len(dirPath) == 0 {
		panic("Empty migration directory path")
	}

	log.Info("Migrating database")

	// Create the ChangeLog table if it doesn't already exist
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ChangeLog (filename VARCHAR(255) NOT NULL)"); err != nil {
		panic(err)
	}

	// Gets all of the migration files found in dirPath.
	fileInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range fileInfo {
		// Determine whether the current migration file has been applied.
		var count int
		row := db.QueryRow("SELECT COUNT(*) FROM ChangeLog WHERE filename = $1", file.Name())
		if err := row.Scan(&count); err != nil {
			panic(err)
		}

		// If not, apply it, and add the migration info to the ChangeLog
		if count == 0 {
			log.Info("Applying migration", log.String("filename", file.Name()))

			// Read the query from the file.
			query, err := ioutil.ReadFile(path.Join(dirPath, file.Name()))
			if err != nil {
				panic(err)
			}

			// Apply the migration.
			if _, err := db.Exec(string(query)); err != nil {
				panic(err)
			}

			// Update the ChangeLog
			if _, err := db.Exec("INSERT INTO ChangeLog (filename) VALUES ($1)", file.Name()); err != nil {
				panic(err)
			}
		}
	}
}
