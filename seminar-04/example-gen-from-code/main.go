package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "example-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Example API
// @version 1.1
// @description Минимальный пример сервера с OpenAPI и разными кодами ошибок
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// API ручки
	r.HandleFunc("/ping", Ping).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")

	// Swagger UI
	r.PathPrefix("/swagger/").
		Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", r)
}

// Ping godoc
// @Summary Проверка сервера
// @Description Возвращает pong
// @Tags system
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetUser godoc
// @Summary Получить пользователя
// @Description Возвращает пользователя по ID или ошибку
// @Tags users
// @Param id path int true "ID пользователя"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse "Неверный ID"
// @Failure 404 {object} ErrorResponse "Пользователь не найден"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	if id == 0 {
		respondWithError(w, http.StatusNotFound, "Пользователь не найден")
		return
	}

	if id == -1 {
		respondWithError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		return
	}

	user := User{
		ID:   id,
		Name: "John",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:    code,
		Message: message,
	})
}
