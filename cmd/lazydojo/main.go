package main

import (
	"database/sql" // Package for SQL database interactions
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kartikm7/lazydojo/internals/models"
	"github.com/kartikm7/lazydojo/pkg/db"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type App struct {
	db *sql.DB
}

func main() {
	db := db.New("./lazydojo.db")
	app := App{db}
	app.RunApp()
}

func (app *App) RunApp() {
	p := tea.NewProgram(models.InitRootModel(app.db), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("shit went south: %s", err)
	}
}
