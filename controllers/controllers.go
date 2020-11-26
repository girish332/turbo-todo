package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/girish332/turbo-todo/dao"

	"github.com/girish332/turbo-todo/model"
	"github.com/gorilla/mux"
)

// Home Function to check if api is working or not
func Home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"isAlive": true}`)
}

// CreateTodo func to create a todo
func CreateTodo(w http.ResponseWriter, r *http.Request) { //Add validations

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content type application/json but got '%s'", ct)))
		return
	}

	var t model.TodoModel
	err = json.Unmarshal(bodyBytes, &t)
	if len(t.Title) == 0 {
		http.Error(w, "Title not entered", http.StatusBadRequest)
	}
	fmt.Println(err)
	fmt.Println(t.Completed)
	t.ID = rand.Intn(100000)
	insertStatement := `INSERT INTO todo (ID, Title, Completed) Values ($1, $2, $3);`
	_, err = dao.DB.Exec(insertStatement, t.ID, t.Title, t.Completed)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error())) //Should return json
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"CreateTodo": true}`)

}

// GetTodos function to get all the todos present in the database
func GetTodos(w http.ResponseWriter, r *http.Request) {

	var todoSlice []model.TodoModel
	getStatement := "select * from todo"
	data, err := dao.DB.Query(getStatement)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for data.Next() {

		var t model.TodoModel
		err = data.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			fmt.Sprintf("Error in data")
			return
		}

		todoSlice = append(todoSlice, t)
	}

	jsonBytes, err := json.Marshal(todoSlice)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

// UpdateTodo Handler to update todo task need to send the updating id via the body
func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content type application/json but got '%s'", ct)))
		return
	}

	var t model.TodoModel
	err = json.Unmarshal(bodyBytes, &t)
	// params := mux.Vars(r)
	// id, _ := params["ID"]
	id := t.ID
	// t.ID = rand.Intn(100000)
	insertStatement := `UPDATE todo SET COMPLETED = $1 WHERE ID = $2;`
	res, err := dao.DB.Exec(insertStatement, t.Completed, id)
	// fmt.Println(t.Completed)
	// fmt.Println(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"updateTodo": true}`)
	// io.WriteString(w, `{"RowsUpdated": count}`)

}

//DeleteTodo func to remove the object from the db
func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	deleteQuery := `DELETE FROM TODO WHERE id = $1;`
	res, err := dao.DB.Exec(deleteQuery, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"DeleteTodo": true}`)

}

// GetTodo ...
func GetTodo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	selectQuery := `SELECT * FROM TODO WHERE id = $1;`
	res, err := dao.DB.Query(selectQuery, id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var t model.TodoModel
	err = res.Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		log.Fatalf("Error in data")
		return
	}

	jsonBytes, err := json.Marshal(t)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)

}
