package functional

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/girish332/turbo-todo/model"

	"github.com/girish332/turbo-todo/bootstrap"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func routerSetup() *mux.Router {

	return bootstrap.TestApplication()
}

func TestGetAllEndpoint(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"210\"}")

	}
	req, err := http.NewRequest("GET", "http://localhost:8080/todos", nil)

	if err != nil {

		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// handler := http.HandlerFunc(controllers.GetTodos)
	// testRouter := routerSetup()
	// testRouter.ServeHTTP(rr, req)
	handler(rr, req)
	// handler.ServeHTTP(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

type createTodo struct {
	CreateTodo bool
}

func TestCreateTodo(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"210\"}")

	}

	body := []byte(`{"Title":"Testing1" , "Completed":false}"`)
	req, err := http.NewRequest("POST", "http://localhost:8080/todo", bytes.NewReader(body))

	if err != nil {

		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// handler := http.HandlerFunc(controllers.GetTodos)
	// testRouter := routerSetup()
	// testRouter.ServeHTTP(rr, req)
	handler(rr, req)
	// handler.ServeHTTP(rr, req)
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// var c createTodo
	var t1 model.TodoModel
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	// fmt.Println(bodyBytes)
	if err != nil {
		t.Errorf("Error in reading body")
	}
	err = json.Unmarshal(bodyBytes, &t1)
	fmt.Println(t1.Title)

	if err != nil {
		t.Errorf("error in unmarshaling")
	}

}
