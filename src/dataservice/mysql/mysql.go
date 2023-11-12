package mysql

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenMySqlConnection() (*sqlx.DB, error) {
	mysqlConfig := fmt.Sprintf("%s:%s@%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
	)

	db, err := sqlx.Connect("mysql", mysqlConfig)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetConnMaxIdleTime(3000)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
