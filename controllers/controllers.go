package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/girish332/turbo-todo/utils"

	"github.com/girish332/turbo-todo/dao"

	"github.com/girish332/turbo-todo/model"
	"github.com/gorilla/mux"
)

// Home Function to check if api is working or not
func Home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// io.WriteString(w, `{"isAlive": true}`)
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
	// insertStatement := `INSERT INTO todo (ID, Title, Completed) Values ($1, $2, $3);`
	// _, err = dao.DB.Exec(insertStatement, t.ID, t.Title, t.Completed)

	err = dao.InsertTodo(t)
	if err != nil {

		utils.JSONError(w, err, 400)
		return
	}

	utils.JSONOk(w, "")
	json.NewEncoder(w).Encode(t)
	// io.WriteString(w, `{"CreateTodo": true}`)

}

// GetTodos function to get all the todos present in the database
func GetTodos(w http.ResponseWriter, r *http.Request) {

	var todoSlice []model.TodoModel

	todoSlice, err := dao.GetAll()
	if err != nil {

		utils.JSONError(w, err, 500)
		return
	}

	_, err = json.Marshal(todoSlice)

	if err != nil {

		utils.JSONError(w, err, 500)
		return
	}

	utils.JSONOk(w, todoSlice)

}

// UpdateTodo Handler to update todo as completed
func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {

		utils.JSONError(w, err, 400)
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {

		utils.JSONError(w, err, 415)
		w.Write([]byte(fmt.Sprintf("Need content type application/json but got '%s'", ct)))
		return
	}

	var t model.TodoModel
	err = json.Unmarshal(bodyBytes, &t)

	id := t.ID
	completed := t.Completed
	count, err := dao.Update(id, completed)
	if err != nil {

		utils.JSONError(w, err, 400)
		return
	}

	fmt.Println(count)

	utils.JSONOk(w, "")
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

	ct := r.Header.Get("content-type")

	if ct != "application/json" {

		utils.JSONError(w, err, 415)
		w.Write([]byte(fmt.Sprintf("Need content type application/json but got '%s'", ct)))
		return
	}

	// deleteQuery := `DELETE FROM TODO WHERE id = $1;`
	// res, err := dao.DB.Exec(deleteQuery, id)

	err = dao.Delete(id)

	if err != nil {
		utils.JSONError(w, err, 400)
		return
	}

	utils.JSONOk(w, "")
	// w.WriteHeader(http.StatusOK)
	// io.WriteString(w, `{"DeleteTodo": true}`)

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
	var t model.TodoModel
	t, err = dao.GetOne(id)

	if err != nil {

		utils.JSONError(w, err, 404)
		fmt.Println(err)
		return
	}

	utils.JSONOk(w, t)

}
