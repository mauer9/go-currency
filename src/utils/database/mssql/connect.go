package mssql

import (
	"database/sql"
	"fmt"
	"net/url"
	"test/src/utils/errors"

	_ "github.com/denisenkom/go-mssqldb"
)

// InitConnect
// initialize Microsoft SQL connection
func InitConnect(host, user, pass, dbName string, port int) (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", dbName)

	uri := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(user, pass),
		Host:     fmt.Sprintf("%s:%d", host, port),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", uri.String())
	if err != nil {
		return nil, errors.DBConnectError.SetDevMessage(err.Error())
	}

	return db, nil
}
