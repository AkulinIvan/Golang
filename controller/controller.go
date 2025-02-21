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
			"task":    task,
		})

}

// Изменяем/обновляем задачу

func UpdateTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task *model.Tasks

	db.DB.Find(&task, taskId)

	if task.Title == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}

	var updateTask model.Tasks
	err := c.BodyParser(&updateTask)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": err,
			})
	}
	if updateTask.Title == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Tasks title is required",
			})
	}
	if updateTask.Description == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Tasks description is required",
			})
	}
	if updateTask.Status == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Tasks status is required",
			})
	}
	//Сохранение задачи в бд

	task.Title = updateTask.Title
	task.Description = updateTask.Description
	task.Status = updateTask.Status
	db.DB.Save(&task)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task changed successfully",
		"task":    task,
	})
}

// Удаляем задачу
func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task model.Tasks

	db.DB.First(&task, taskId)

	db.DB.Delete(&task)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task was deleted successfuly",
	})
}
