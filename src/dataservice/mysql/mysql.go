package mysql

import (
	"fmt"
	"golang-backend-microservice/container/logger"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenMySqlConnection() *sqlx.DB {
	mysqlConfig := fmt.Sprintf("%s:%s@%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASS"),
		os.Getenv("MYSQL_HOST"),
	)

	db, err := sqlx.Connect("mysql", mysqlConfig)
	if err != nil {
		logger.Error("Error connecting to MySQL: ", err)
		log.Fatalf("Error connecting to MySQL: %s\n", err)
		return nil
	}
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetConnMaxIdleTime(3000)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		logger.Error("Error pinging to MySQL: ", err)
		log.Fatalf("Error pinging to MySQL: %s\n", err)
		return nil
	}

	return db
}
