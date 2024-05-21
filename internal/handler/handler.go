package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/internal/storage"
	"github.com/Yandex-Practicum/go-rest-api-homework/internal/task"
	"github.com/go-chi/chi"
)

// Обработчик для получения всех задач
func Tasks(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(storage.Tasks)
	if err != nil {
		log.Println("can't marshal json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		log.Println("can't write response:", err)
	}
}

// Обработчик для получения задачи по ID
func Task(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	task, ok := storage.Tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusNoContent)
		return
	}

	resp, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		log.Println("can't write response:", err)
	}
}

// Обработчик для отправки задачи на сервер
func PostTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage.Tasks[task.ID] = task

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Обработчик удаления задачи по ID
func DelTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := storage.Tasks[id]
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusNoContent)
		return
	}

	delete(storage.Tasks, id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
