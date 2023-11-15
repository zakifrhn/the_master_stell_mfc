package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Postgres_db() (*sqlx.DB, error) {

	host := viper.GetString("database.host")
	dbname := viper.GetString("database.name")
	port := viper.GetString("database.port")
	password := viper.GetString("database.pass")
	user := viper.GetString("database.user")

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sqlx.Connect("postgres", config)
}
