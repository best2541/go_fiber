package controllers

import (
	"fmt"
	"go_fiber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Try(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "ok",
	})
}

func test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hi",
	})
}
func GetAllFoods(c *fiber.Ctx) error {
	return c.JSON(models.Foods)
}
func GetFood(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32) // Base 10, 32-bit size
	if err != nil {
		panic(err)
	}
	for _, food := range models.Foods {
		fmt.Println(id)
		fmt.Println(food)
		if food.ID == uint(idUint) {
			return c.JSON(food)
		}
	}
	return nil
}
func InsertFood(c *fiber.Ctx) error {
	var food models.Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}
	models.Foods = append(models.Foods, food)
	return c.JSON(food)
}
func EditFoods(c *fiber.Ctx) error {
	var editFood models.EditFood
	if err := c.BodyParser(&editFood); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32) // Base 10, 32-bit size
	if err != nil {
		panic(err)
	}

	for i := range models.Foods {
		if models.Foods[i].ID == uint(idUint) {
			models.Foods[i].Name = editFood.Name
			models.Foods[i].Price = editFood.Price
			return c.SendString("Edited")
		}
	}
	return nil
}

func DeleteFood(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)
	}
	for i, food := range models.Foods {
		if food.ID == uint(idUint) {
			models.Foods = append(models.Foods[:i], models.Foods[i+1:]...)
			return c.SendString("deleted")
		}
	}
	return nil
}

func GetInfo(c *fiber.Ctx) error { // JSON
	return c.JSON(fiber.Map{
		"msg":     "hello world 🚀",
		"go":      "fiber 🥦",
		"boolean": true,
		"number":  1234,
	})
}
