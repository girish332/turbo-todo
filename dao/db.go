package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/girish332/turbo-todo/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DB ...
var DB *sql.DB

// DatabaseInit func to initialize connection to DB
func DatabaseInit() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbDsn := os.Getenv("dbDsn")
	DB, err = sql.Open("postgres", dbDsn)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
}

// GetAll ....
func GetAll() (slice []model.TodoModel, err error) {

	getStatement := "select * from todo"
	data, err := DB.Query(getStatement)

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
func InsertTodo(t model.TodoModel) (err error) {

	insertStatement := `INSERT INTO todo (ID, Title, Completed) Values ($1, $2, $3);`
	_, err = DB.Exec(insertStatement, t.ID, t.Title, t.Completed)

	if err != nil {
		return err
	}

	return nil

}

// GetOne function to get one todo from db
func GetOne(id int) (t model.TodoModel, err error) {

	var t1 model.TodoModel
	selectQuery := `SELECT * FROM TODO WHERE id = $1;`
	res, err := DB.Query(selectQuery, id)

	if err != nil {
		return t1, err
	}

	for res.Next() {
		err = res.Scan(&t1.ID, &t1.Title, &t1.Completed)
		if err != nil {
			return t1, err
		}
	}

	return t1, nil
}

// Update ...
func Update(id int, completed bool) (count int64, err error) {

	insertStatement := `UPDATE todo SET COMPLETED = $1 WHERE ID = $2;`
	res, err := DB.Exec(insertStatement, completed, id)

	if err != nil {
		return 0, err
	}

	c, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return c, nil

}

// Delete ...
func Delete(id int) (err error) {

	deleteQuery := `DELETE FROM TODO WHERE id = $1;`
	_, err = DB.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	return nil
}
