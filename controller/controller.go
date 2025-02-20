package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Создаем задачу
func CreateTask(c *fiber.Ctx) error {
	return nil
}

// Выводим список задач
func ListOfTasks(c *fiber.Ctx) error {
	return c.SendString("list of tasks")

}

// Изменяем/обновляем задачу
func UpdateTask(c *fiber.Ctx) error {
	return nil
}

// Удаляем задачу
func DeleteTask(c *fiber.Ctx) error {
	return nil
}
