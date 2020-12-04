package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/girish332/turbo-todo/controllers"

	"github.com/girish332/turbo-todo/dao"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func InitDatabase() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
		fmt.Println(err)
	}
	dbDsn := os.Getenv("dbDsn")
	DB, err := sql.Open("postgres", dbDsn)

	return DB
}

func main() {

	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
		fmt.Println(err)
	}
	dbDsn := os.Getenv("dbDsn")
	DB, err := sql.Open("postgres", dbDsn)

	env := &controllers.Controller{
		Todo: dao.DataBase{DB: DB},
	}

	fmt.Println("Connected to DB")

	router := mux.NewRouter()
	// router.HandleFunc("/home", controllers.Home).Methods("GET")
	router.HandleFunc("/todo", env.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", env.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", env.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", env.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}", env.GetTodo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	// bootstrap.BootApplication()

}
