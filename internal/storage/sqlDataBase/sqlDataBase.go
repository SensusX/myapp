package sqlDataBase

import (
	"database/sql"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SqlDataBase struct {
	db             *sql.DB
	Driver         string
	DataSourceName string
}

func (sqlDataBase *SqlDataBase) Open(log slog.Logger) {

	log.Info("Open SQL Data Base...")
	var err error
	sqlDataBase.db, err = sql.Open(sqlDataBase.Driver, sqlDataBase.DataSourceName)
	if err != nil {
		log.Info("ERROR OPEN SQLDATABASE: " + err.Error())
		return
	}

	stm, err := sqlDataBase.db.Prepare(`CREATE TABLE IF NOT EXISTS Users (
    id integer PRIMARY KEY,
    LastName varchar(255),
    FirstName varchar(255)
    );`)
	if err != nil {
		log.Debug("ERROR PREPARE SQLDATABASE: " + err.Error())
		return
	}

	_, err = stm.Exec()
	if err != nil {
		log.Debug("ERROR Exec SQLDATABASE")
		return
	}
}
