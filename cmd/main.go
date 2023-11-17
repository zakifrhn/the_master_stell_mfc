package main

import (
	"fmt"
	"log"
	"test-hiring/internal/routers"
	"test-hiring/pkg"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := pkg.Postgres_db()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	router := routers.Routers(db)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

//?create table
//migrate -path ./migrations -database "postgresql://postgres:Fazztrak2023@localhost/go_test_hiring?port=5432&sslmode=disable&search_path=public" -verbose up

//?delete table
//migrate -path ./migrations -database "postgresql://postgres:Fazztrak2023@localhost/go_test_hiring?port=5432&sslmode=disable&search_path=public" -verbose down
