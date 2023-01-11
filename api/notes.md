go install github.com/cosmtrek/air@latest

`Instala in hot reloading para Go `

go get -u gorm.io/gorm
`instala GORM`

go get -u gorm.io/driver/postgres
`instala el Driver`


example using .env 

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// connect to the database
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	
	fmt.Println("Connection to the database is successful")
	
	// do database operations
	...
}
