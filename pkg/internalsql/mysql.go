package internalsql

import (
	"database/sql"
	"fmt"
	"go-tweets/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(cfg *config.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword,
	// dataSourceName := fmt.Sprintf("root:12345@127.0.0.1:3306/go-tweets?parseTime=true")

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}
	log.Panicln("database connected")
	return db, nil
}
