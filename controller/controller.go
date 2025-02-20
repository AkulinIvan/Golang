package controller

import (
	"time"

	db "github.com/AkulinIvan/Golang/config"
	"github.com/AkulinIvan/Golang/model"
	"github.com/gofiber/fiber/v2"
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
	if data["Title"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Tasks title is required",
			})
	}
	if data["Description"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Tasks description is required",
			})
	}
	//Сохранение задачи в бд

	task := model.Tasks{
		Title:       data["Title"],
		Description: data["Description"],
		Status:      model.NewStatus,
		Created_at:  time.Time{},
		Updated_at:  time.Time{},
	}
	db.DB.Create(&task)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task added successfully",
		"data":    task,
	})
}

// Выводим список задач
func ListOfTasks(c *fiber.Ctx) error {
	var task []model.Tasks

	db.DB.Find(&task)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Tasks list api",
			"data":    task,
		})

}

// Изменяем/обновляем задачу
func UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task model.Tasks

	db.DB.Find(&task, "id=?", taskId)
	//валидация для проверки id задачи

	if task.Title != "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}

	var updateTask model.Tasks
	err := c.BodyParser(&updateTask)
	if err != nil {
		return err
	}
	if updateTask.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Task title is required",
		})
	}

	task.Title = updateTask.Title
	db.DB.Save(&task)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    task,
	})

}

// Удаляем задачу
func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task model.Tasks
	result := db.DB.Delete(&task, "id = ?", taskId)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No task with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
