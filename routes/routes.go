package routes

import (
	Controller "github.com/AkulinIvan/Golang/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Создаем новую задачу
	app.Post("/tasks/", Controller.CreateTask)
	//Получаем список всех задач
	app.Get("/tasks/", Controller.ListOfTasks)

	//Обновляем/изменяем задачу
	app.Put("/tasks/:Id", Controller.UpdateTask)
	//Удаляем задачу
	app.Delete("/tasks/:Id", Controller.DeleteTask)
	//Слушаем приложение на порту 3000
	app.Listen(":3000")
}
