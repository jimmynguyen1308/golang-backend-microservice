package mysql

import (
	"encoding/json"
	"golang-backend-microservice/container/log"
	"golang-backend-microservice/model"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go/micro"
)

func SelectRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := buildSelectQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}

		switch data.Table {
		case "book":
			var records []model.Book
			for results.Next() {
				var record model.Book
				if err := results.StructScan(&record); err != nil {
					log.Error(err.Error())
					continue
				}
				records = append(records, record)
			}
			log.Info("Output: %v", records)
			// TODO: send Nats response
		default:
			log.Error(log.ErrMySqlUnknwonTable, data.Table)
			// TODO: send Nats response
		}
		results.Rows.Close()
	}
}

func InsertRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := buildInsertQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}

func UpdateRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := buildUpdateQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}

func DeleteRecord(db *sqlx.DB) micro.HandlerFunc {
	return func(req micro.Request) {
		var data model.MySqlReqArgs
		json.Unmarshal(req.Data(), &data)
		query, args, err := buildDeleteQuery(&data)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results, err := db.Queryx(query, args...)
		if err != nil {
			log.Error(err.Error())
			// TODO: send Nats response
			return
		}
		results.Rows.Close()
	}
}
