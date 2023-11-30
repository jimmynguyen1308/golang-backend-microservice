package mysql

import (
	"fmt"
	"golang-backend-microservice/container/log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	User string // MySQL username
	Pass string // MySQL password
	Host string // MySQL database name
}

func (c Connection) Open() *sqlx.DB {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@/%s", c.User, c.Pass, c.Host))
	if err != nil {
		log.Error(log.ErrMySqlConnect, err.Error())
		return nil
	}
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetConnMaxIdleTime(3000)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Error(log.ErrMySqlConnect, err.Error())
		return nil
	}

	return db
}
