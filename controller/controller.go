package controller

import (
	"github.com/gofiber/fiber/v3"
)

// Создаем задачу
func CreateTask(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			})
	}

	if data["title"] = ""{
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Title is required",
			}
		)
	}
}

// Выводим список задач
func ListOfTasks(c *fiber.Ctx) error {
	c.SendString("tasks list")
}

// Изменяем/обновляем задачу
func UpdateTask(c *fiber.Ctx) error {
	return nil
}

// Удаляем задачу
func DeleteTask(c *fiber.Ctx) error {
	return nil
}
