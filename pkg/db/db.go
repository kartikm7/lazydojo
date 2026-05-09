/* Package db defines a bunch of basic db related methods
* **/
package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/charmbracelet/log"
)

type Task struct {
	ID        int
	Task      string
	Completed bool
}

type Values struct {
	Completed *bool
	Task      *string
}

func New(src string) *sql.DB {
	db, err := sql.Open("sqlite3", src)
	if err != nil {
		log.Fatalf("shit went wrong with the db: %s", err)
	}

	// we need to create the table too dawgs
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		task TEXT,
		completed BOOLEAN
		)
		`)
	if err != nil {
		log.Fatalf("shit went wrong with the query: %s", err)
	}
	return db
}

// Add allows us to add a task to the sqlite3 db
// it also predefines the task completion as false and time spent as empty
func Add(db *sql.DB, task string) error {
	_, err := db.Exec("INSERT INTO tasks(task, completed) VALUES(?,?)", task, false)
	return err
}

func Delete(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func Update(db *sql.DB, id int, values Values) error {
	v := reflect.ValueOf(values)
	for i := 0; i < v.NumField(); i++ {
		column := v.Type().Field(i).Name
		value := v.Field(i)
		if value.Interface() != nil {
			// man this shit is the coolest thing ever, so damn handy
			if value.IsZero() {
				slog.Info("[Database] Not a valid value continuing to next iteration")
				continue
			}
			value := value.Elem()
			convertedString := fmt.Sprintf("%v", value)
			query := fmt.Sprintf("UPDATE tasks SET %v = ? WHERE id = ?", strings.ToLower(column))
			slog.Info("[Database] Running the update method", "query", query)
			// Now, the problem with the above log is that we don't print the passed values
			// But, I think the db.Exec should be having some security placeholders - still does not feel right
			result, err := db.Exec(query, convertedString, id)
			if err != nil {
				slog.Info("[Database] Fuck something went wrong", "error", err.Error())
			} else {
				slog.Info("[Database] Successfully ran the query", "result", result)
			}
		}
	}
	return nil
}

func ListEverything(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	tasks := []Task{}
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Task, &task.Completed); err != nil {
			log.Printf("Something went wrong %s", err)
			return nil, err
		}
		defer rows.Close()
		tasks = append(tasks, task)
	}

	return tasks, err
}
