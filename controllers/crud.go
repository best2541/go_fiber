package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// // User model
// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

func GetUsers(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT user_id FROM USERS")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	fmt.Print(rows)
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {
		var user_id string
		err := rows.Scan(&user_id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to scan user data"})
		}
		users = append(users, map[string]interface{}{
			"user_id": user_id,
		})
	}

	return c.JSON(users)
}

// Handler to create a new user
func CreateUser(c *fiber.Ctx, db *sql.DB) error {
	type User struct {
		User_id string `json:"user_id"`
	}

	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	_, err := db.Exec("INSERT INTO USERS (user_id) VALUES (?)", user.User_id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}
