package routes

import (
	controller "github.com/AkulinIvan/go-test/controller"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	// Создаем новый экземпляр Fiber
	

	// Создаем новую задачу
	app.Post("/tasks/", controller.CreateTask)
	//Получаем список всех задач
	app.Get("/tasks", controller.ListOfTasks)
}

// 	controller.CreateTask})

// //Обновляем/изменяем задачу
// app.Put("/tasks/:id", controller.UpdateTask)
// //Удаляем задачу
// app.Delete("/tasks/:id", controller.DeleteTask)
//Слушаем приложение на порту 3000
// app.Listen(":3000")
