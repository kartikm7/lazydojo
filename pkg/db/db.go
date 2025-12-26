/* Package db defines a bunch of basic db related methods
* **/
package db

import (
	"database/sql"
	"log"
)

func New(src string) *sql.DB {
	db, err := sql.Open("sqlite3", src)
	if err != nil {
		log.Fatalf("shit went wrong with the db: %s", err)
	}

	// we need to create the table too dawgs
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		task TEXT
		)
		`)
	if err != nil {
		log.Fatalf("shit went wrong with the query: %s", err)
	}
	return db
}

func Add(db *sql.DB, task string) error {
	_, err := db.Exec("INSERT INTO tasks(task) VALUES(?)", task)
	return err
}

func Delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
