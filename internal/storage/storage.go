package storage

import (
	"errors"
	"log/slog"

	. "github.com/SensusX/myapp/internal/storage/sqlDataBase"
)

type Storage interface {
	Open(log slog.Logger)
}

func New(storageType string, driver string, log slog.Logger) (Storage, error) {

	log = *log.With("func", "storage.New")
	switch storageType {
	case "SQL":
		log.Info("Chosen SQL Data Base")
		var sqlDataBase SqlDataBase
		sqlDataBase.Driver = driver
		sqlDataBase.DataSourceName = "user=myappuser password=myappuser host=localhost dbname=myapp port=5432"
		return &sqlDataBase, nil
	}
	return nil, errors.New("ERROR IN NEW STORAGE")
}
