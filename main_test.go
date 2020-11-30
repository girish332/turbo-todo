package main

import (
	"log"
	"os"
	"testing"

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
