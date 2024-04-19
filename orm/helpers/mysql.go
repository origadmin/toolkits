package helpers

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

const (
	mysqlDatabaseCreateFmt = "CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET '%s' DEFAULT COLLATE '%s';"
	mysqlCharSet           = "utf8mb4"
	mysqlCollate           = "utf8mb4_general_ci"
)

// CreateMySQLDatabase creates a MySQL database with the given DSN.
func CreateMySQLDatabase(dsn string) error {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr))
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Failed to close database: %v\n", err)
		}
	}(db)

	charset := mysqlCharSet
	if cs, ok := cfg.Params["charset"]; ok {
		charset = cs
	}
	collate := mysqlCollate
	if cfg.Collation != "" {
		collate = cfg.Collation
	}

	query := fmt.Sprintf(mysqlDatabaseCreateFmt, cfg.DBName, charset, collate)
	_, err = db.Exec(query)
	return err
}
