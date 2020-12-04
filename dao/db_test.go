package dao

import (
	"fmt"
	"testing"

	"github.com/girish332/turbo-todo/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// SELECT id,name,email,phone FROM users WHERE id = \\?

// SELECT id,name,email,phone FROM users WHERE id = ?

func TestGetOne(t *testing.T) {

	db, mock, err := sqlmock.New()
	repo := &DataBase{db}

	query := `SELECT id,title,completed FROM TODO WHERE id \= \$1`
	rows := sqlmock.NewRows([]string{"Id", "Title", "Completed"}).AddRow(1, "test1", false)

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	todo, err := repo.GetOne(1)

	if todo.ID != 1 {
		t.Errorf("Wrong id")
	}
	// assert.NotNil(t, todo)
	assert.NoError(t, err)

}

func TestCreateOne(t *testing.T) {

	db, mock, err := sqlmock.New()
	repo := &DataBase{db}

	query := "INSERT INTO todo \\(id, title, completed\\) VALUES \\(\\?, \\?, \\?\\)"
	// rows := sqlmock.NewRows([]string{"Id", "Title", "Completed"}).AddRow(1, "test1", false)

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1, "test2", false).WillReturnResult(nil)
	obj := model.TodoModel{
		ID:        1,
		Title:     "test2",
		Completed: false,
	}
	err = repo.InsertTodo(obj)
	// assert.NotNil(t, todo)
	assert.NoError(t, err)

}

func TestGetAll(t *testing.T) {

	db, mock, err := sqlmock.New()
	repo := &DataBase{db}

	query := `SELECT id,title,completed FROM TODO`
	rows := sqlmock.NewRows([]string{"Id", "Title", "Completed"}).AddRow(1, "test1", false)

	mock.ExpectQuery(query).WillReturnRows(rows)

	_, err = repo.GetAll()

	// assert.NotNil(t, todo)
	assert.NoError(t, err)

}

func TestUpdate(t *testing.T) {

	db, mock, err := sqlmock.New()
	repo := &DataBase{db}

	query := `UPDATE todo SET completed \= \$1 WHERE id \= \$2`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(0, 1))

	c, err := repo.Update(1, true)
	fmt.Println(c)
	if c != 1 {
		t.Errorf("value not updated")
	}
	assert.Error(t, err)
}
