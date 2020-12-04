package dao

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/girish332/turbo-todo/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB ...
// var DB *sql.DB

type DataBase struct {
	DB *sql.DB
}

type databaseServiceInterface interface {
	GetAll() (slice []model.TodoModel, err error)
	InsertTodo(t model.TodoModel) (err error)
	GetOne(id int) (t model.TodoModel, err error)
	Update(id int, completed bool) (count int64, err error)
	Delete(id int) (err error)
}

func NewMockDatabase(db *sql.DB) databaseServiceInterface {
	return &DataBase{DB: db}
}

type todoModel struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// DatabaseInit func to initialize connection to DB
func DatabaseInit() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
		fmt.Println(err)
	}
	dbDsn := os.Getenv("dbDsn")
	DB, err := sql.Open("postgres", dbDsn)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
	return DB
}

// GetAll ....
func (d DataBase) GetAll() (slice []model.TodoModel, err error) {

	getStatement := "SELECT id,title,completed FROM TODO"
	data, err := d.DB.Query(getStatement)

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var todoSlice []model.TodoModel

	for data.Next() {

		var t model.TodoModel
		err = data.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			return nil, err
		}

		todoSlice = append(todoSlice, t)
	}

	return todoSlice, nil

}

//InsertTodo func to insert into db
func (d DataBase) InsertTodo(t model.TodoModel) (err error) {

	insertStatement := "INSERT INTO todo (id, title, completed) VALUES (?, ?, ?)"
	_, err = d.DB.Exec(insertStatement, t.ID, t.Title, t.Completed)

	if err != nil {
		return err
	}

	return nil

}

// GetOne function to get one todo from db
func (d DataBase) GetOne(id int) (t model.TodoModel, err error) {

	var t1 model.TodoModel
	selectQuery := "SELECT id,title,completed FROM TODO WHERE id = $1"
	res := d.DB.QueryRow(selectQuery, id)

	err = res.Scan(&t1.ID, &t1.Title, &t1.Completed)

	if err != nil {
		return t1, err
	}

	return t1, nil
}

// Update ...
func (d DataBase) Update(id int, completed bool) (count int64, err error) {

	insertStatement := `UPDATE todo SET COMPLETED = $1 WHERE ID = $2`
	res, err := d.DB.Exec(insertStatement, completed, id)

	if err != nil {
		return 1, err
	}

	c, err := res.RowsAffected()
	if err != nil {
		return 1, err
	}

	return c, nil

}

// Delete ...
func (d DataBase) Delete(id int) (err error) {

	deleteQuery := `DELETE FROM TODO WHERE id = $1;`
	_, err = d.DB.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	return nil
}
