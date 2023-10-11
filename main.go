package main

import (
	"fmt"
	"library-backend-new1/internal/models"
	"library-backend-new1/internal/routes"

	// "gorm.io/gorm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// func ConnectDb() {
// 	host := "localhost"
// 	port := "5433"
// 	database := "Library_New"
// 	user_name := "postgres"
// 	password := "Singha@12"

// 	db, err := gorm.Open(postgres.Open("postgres://" + user_name + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"))
// 	if err != nil {
// 		fmt.Println(err, " Database Connection Failed")
// 		log.Fatal("connection error: ", err)
// 	} else {
// 		fmt.Println("Successfully connected to the database")
// 	}
// 	db.AutoMigrate(&models.Library{}, &models.User{}, &models.BookInventory{}, &models.RequestEvents{}, &models.IssueRegistry{})
// }

var db *gorm.DB //database
func main() {
	// Initialize database connection
	// host := "localhost"
	// port := "5433"
	// database := "Library_New"
	// user_name := "postgres"
	// password := "Singha@12"
	// dsn := "postgres://" + user_name + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"
	// db, err := gorm.Open(postgres.Open("postgres://" + user_name + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"))

	username := "postgres"
	password := "Singha@12"
	dbName := "Library_0706"
	dbHost := "localhost"
	port := "5433"
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, port, username, dbName, password) //Build connection string
	fmt.Println(dbUri)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	defer db.Close() // Remember to close the connection when you're done with it

	// AutoMigrate will create the necessary tables based on the models
	db.AutoMigrate(&models.Library{}, &models.User{}, &models.BookInventory{}, &models.RequestEvents{}, &models.IssueRegistry{})
	fmt.Println("===========", db)

	// Create a new instance of a Gin router
	r := routes.SetupRouter(db)
	// Start the server
	r.Run(":8080")
}
