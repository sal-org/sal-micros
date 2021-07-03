package database

import (
	"database/sql"
	"fmt"
	"log"
	CONFIG "notification/config"

	_ "github.com/go-sql-driver/mysql" // for mysql driver
)

var db *sql.DB
var err error

// ConnectDatabase - connect to mysql database with given configuration
func ConnectDatabase() {
	db, err = sql.Open("mysql", CONFIG.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
}

// sql wrapper functions

// QueryRowSQL - get single data with defined values
func QueryRowSQL(SQLQuery string, params ...interface{}) string {
	var value string
	db.QueryRow(SQLQuery, params...).Scan(&value)
	return value
}

// ExecuteSQL - execute statement with defined values
func ExecuteSQL(SQLQuery string, params ...interface{}) (sql.Result, error) {
	return db.Exec(SQLQuery, params...)
}

// SelectProcess - execute raw select statement
func SelectProcess(SQLQuery string, params ...interface{}) ([]map[string]string, bool) {

	rows, err := db.Query(SQLQuery, params...)
	if err != nil {
		fmt.Println("SelectProcess", err)
		return []map[string]string{}, false // default
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("SelectProcess", err)
		return []map[string]string{}, false // default
	}

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols))
	data := []map[string]string{}
	rest := map[string]string{}
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		rest = map[string]string{}
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("SelectProcess", err)
			return []map[string]string{}, false // default
		}

		for i, raw := range rawResult {
			if raw == nil {
				rest[cols[i]] = ""
			} else {
				rest[cols[i]] = string(raw)
			}
		}

		data = append(data, rest)
	}
	return data, true
}
