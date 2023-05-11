package mysql1

import (
	"database/sql"
	"fmt"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB db
var DB *sql.DB

// Init db
func Init(host string, port int, user string, password string, dbName string) (err error) {
	DB, err = Open(host, port, user, password, dbName)
	return
}

// Open mysql
func Open(host string, port int, user string, password string, dbName string) (db *sql.DB, err error) {
	defer func() {
		if rev := recover(); rev != nil {
			err = fmt.Errorf("mysql: Open err: %v", rev)
		}
	}()

	sourceStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FShanghai", user, password, host, port, dbName)
	if db, err = sql.Open("mysql", sourceStr); err != nil {
		err = fmt.Errorf("Open: %v", err)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)

	if err = db.Ping(); err != nil {
		err = fmt.Errorf("ping: %v", err)
		return
	}
	return
}
