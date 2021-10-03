package database

import (
	"database/sql"
	"fmt"
	"log"
	CONFIG "refund/config"

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

// UpdateSQL - update data with defined values
func UpdateSQL(tableName string, params map[string]string, body map[string]string) bool {
	args := []interface{}{}

	if len(body) == 0 {
		return false
	}
	SQLQuery := "update `" + tableName + "` set "

	init := false
	for key, val := range body {
		if init {
			SQLQuery += ","
		}
		SQLQuery += "`" + key + "` = ? "
		args = append(args, val)
		init = true
	}

	SQLQuery += " where "
	init = false
	for key, val := range params {
		if init {
			SQLQuery += " and "
		}
		SQLQuery += "`" + key + "` = ? "
		args = append(args, val)
		init = true
	}

	_, err = db.Exec(SQLQuery, args...)
	if err != nil {
		fmt.Println("UpdateSQL", err)
		return false // default
	}
	return true
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
