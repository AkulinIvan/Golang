package controller

import (
	"strconv"
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
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&task).Count(&count)
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

	if task.Title == "" {
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

	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    task,
	})

}

// Удаляем задачу
func DeleteTask(c *fiber.Ctx) error {
	taskId := c.Params("taskId")
	var task model.Tasks

	db.DB.Where("id=?", taskId).First(&task)

	if task.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Task not found",
		})
	}
	db.DB.Where("id=?", taskId).Delete(&task)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task was deleted successfuly",
	})
}
