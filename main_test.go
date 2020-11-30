package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/girish332/turbo-todo/controllers"
	"github.com/girish332/turbo-todo/dao"
)

func TestMain(m *testing.M) {
	dao.DatabaseInit()

	ensureTableExists()
	code := m.Run()
	clearTable()

	os.Exit(code)

}

func ensureTableExists() {
	if _, err := dao.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS TODO
(
	id INTEGER,
	title VARCHAR,
	completed boolean
    
)`

func clearTable() {
	dao.DB.Exec("DELETE FROM todo")
	dao.DB.Exec("ALTER SEQUENCE todo_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/home", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// if body := response.Body.String(); body != "[]" {
	// 	t.Errorf("Expected an empty array. Got %s", body)
	// }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/home", controllers.Home).Methods("GET")
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "http:localhost/todos/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/todos/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		dao.DB.Exec("INSERT INTO TODO(TITLE, COMPLETED) VALUES($1, $2)", "Todo "+strconv.Itoa(i), false)
	}
}
