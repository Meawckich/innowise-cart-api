package db

import (
	"database/sql"
	"runtime"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitDb() (*sql.DB, error) {

	path := viper.Get("DB_URL").(string)

	pool, err := sql.Open("postgres", path)

	pool.SetMaxIdleConns(1)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(runtime.NumCPU())

	if err != nil {
		return nil, err
	}

	return pool, nil
}
