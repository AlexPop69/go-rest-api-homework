package main

import (
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/internal/handler"

	"github.com/go-chi/chi/v5"
)

const port = ":8080"

func main() {
	r := chi.NewRouter()

	// регистрируем в роутере эндпоинт `/tasks` с методом GET, для которого используется обработчик `getTasks`
	r.Get("/tasks", handler.Tasks)
	// эндпоинт `/tasks{id}` с методом GET, для которого используется обработчик `getTask`
	r.Get("/tasks/{id}", handler.Task)
	// эндпоинт `/tasks` с методом POST, для которого используется обработчик `postTask`
	r.Post("/tasks", handler.PostTask)
	// эндпоинт `/tasks{id}` с методом DELETE, для которого используется обработчик `delTask`
	r.Delete("/tasks/{id}", handler.DelTask)

	log.Println("Server started")

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err.Error())
		return
	}
}
