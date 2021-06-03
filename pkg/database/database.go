package database

import (
	"golang-starter/pkg/config"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

type Config struct {
	Dsn string `json:"dsn"`
}

func Init() error {
	conn, err := sqlx.Connect("pgx", config.GetConfig().DB.Dsn)
	db = conn
	return err
}

func GetDB() *sqlx.DB {
	return db
}

func Close() error {
	return db.Close()
}
