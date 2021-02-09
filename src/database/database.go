package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	sql.Register("mysql", &MySQLDriver{}) // Register the driver? idfk.
}

const ( // Set a fuckload of things we probably don't need,
	username = "CockSucker" // but we have them anyways for the sake of having shit
	password = "FuckMyAss2021"
	hostname = "127.0.0.1:3306"
	dbname   = "users"
)

func fuck() {
	db, err := sql.Open("mysql", username+password+"@"+hostname+"/"+dbname)
	if err != nil {
		panic(err)
	}
	ctx, cancelConn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelConn()
	db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

}
