package bootstrap

import (
	"log"
	"net/http"

	"github.com/girish332/turbo-todo/controllers"
	"github.com/girish332/turbo-todo/dao"

	"github.com/gorilla/mux"
)

// BootApplication ...
func BootApplication() {

	dao.DatabaseInit()
	router := mux.NewRouter()
	router.HandleFunc("/home", controllers.Home).Methods("GET")
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
