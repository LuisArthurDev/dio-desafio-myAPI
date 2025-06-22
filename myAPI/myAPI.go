package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
var nextID int = 1

func loadFromFile() {
	file, err := os.Open("users.json")
	if err != nil {
		log.Println("Arquivo JSON não encontrado. Iniciando lista vazia.")
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		log.Println("Erro ao ler o arquivo JSON:", err)
		return
	}

	for _, user := range users {
		if user.ID >= nextID {
			nextID = user.ID + 1
		}
	}
}

func saveToFile() {
	file, err := os.Create("users.json")
	if err != nil {
		log.Println("Erro ao salvar no arquivo JSON:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(users)
}

// Handlers da API:

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	saveToFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == id {
			var updatedUser User
			err := json.NewDecoder(r.Body).Decode(&updatedUser)
			if err != nil {
				http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
				return
			}

			updatedUser.ID = id
			users[index] = updatedUser
			saveToFile()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			saveToFile()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func main() {
	// Carregar os usuários já existentes do JSON (se houver)
	loadFromFile()

	// Criar o router
	r := mux.NewRouter()

	// Rotas
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Iniciar o servidor
	fmt.Println("Servidor rodando em http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
