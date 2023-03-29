package dberr

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsDuplicateKey(_err error) bool {
	if pgError := _err.(*pgconn.PgError); errors.Is(_err, pgError) {
		switch pgError.Code {
		case "23505":
			return true
		}
	} else if mysqlErr := _err.(*mysql.MySQLError); errors.Is(_err, mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			return true
		}
	}
	return false
}
