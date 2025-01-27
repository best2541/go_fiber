package main

import (
	"database/sql"
	"fmt"
	"go_fiber/controllers"
	"log"

	_ "github.com/go-sql-driver/mysql"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

func main() {
	app := fiber.New()

	// Connect to the MySQL database
	var err error
	dsn := "idewofktlttgg7nz:d1yjz40wualr5w69@tcp(dcrhg4kh56j13bnu.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/aoe1adeab51zt65c"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Ping the database to ensure connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to the MySQL database successfully!")

	app.Get("/", controllers.Try)

	app.Get("/foods", controllers.GetAllFoods)

	app.Get("/foods/:id", controllers.GetFood)

	app.Post("/foods", controllers.InsertFood)

	app.Post("/foods/:id", controllers.EditFoods)

	app.Delete("/foods/:id", controllers.DeleteFood)

	app.Get("/info", controllers.GetInfo)

	app.Post("/login", controllers.Login)

	app.Use("/protected", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: controllers.JwtSecret},
	}))

	app.Get("/protected", controllers.Protected)

	app.Get("/users", func(c *fiber.Ctx) error {
		return controllers.GetUsers(c, db)
	})

	app.Listen(":3000")
}
