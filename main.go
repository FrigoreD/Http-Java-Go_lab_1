package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
*
Псевдо база
*/
var (
	db = map[string]string{
		"user1":       "pass1",
		"superuser":   "superpass",
		"otheruser":   "otheruser",
		"anotheruser": "anotheruser",
	}
)

/*
*
Точка входа
*/
func main() {
	http.HandleFunc("/", loginPageHandler)
	http.HandleFunc("/login", loginRequestHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/adduser", addUserRequestHandler)
	http.HandleFunc("/adduserpage", addUserPageHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/*
*
Странмица login.html
*/
func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

/*
*
Запрос с формы на авторизацию по введённому логину/паролю с редиректом на /time или выводом ошибки на экран
*/
func loginRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if isValidCredentials(username, password) {
		http.Redirect(w, r, "/time", http.StatusSeeOther)
		return
	}

	// Если связка логин/пароль не верны, выводим ошибку
	http.Error(w, "Неверные данные", http.StatusUnauthorized)
}

/*
*
Страница с выводом текущего time
*/
func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Текущее время: %s .\n", time.Now())
}

/*
*
Метод добавление пользователя в псевдо базу
*/
func addUserRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не доступен", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if confirmPassword != password {
		http.Error(w, "Пароли не совпадают", http.StatusBadRequest)
		return
	}

	if username != "" {
		db[username] = password
		fmt.Fprintf(w, "Пользователь: %s добавлен в базу данных", username)
		return
	}

	http.Error(w, "Непредвиденная ошибка", http.StatusBadRequest)
}

/*
*
Метод проверки связки и сопоставлением с данными из бд
*/
func isValidCredentials(username, password string) bool {
	storedPassword, ok := db[username]
	if !ok {
		return false
	}
	return storedPassword == password
}

/*
*
Страница с формой регистрации пользователя
*/
func addUserPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "add_user_page.html")
}
