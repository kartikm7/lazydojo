package main

import (
	"database/sql" // Package for SQL database interactions
	"log/slog"
	"os"

	"github.com/charmbracelet/log"

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
	f := setupLogging()
	defer f.Close()
	if _, err := p.Run(); err != nil {
		log.Fatalf("shit went south: %s", err)
	}
}

func setupLogging() *os.File {
	f, err := os.OpenFile("./app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("shit went south: %s", err)
	}

	handler := slog.NewTextHandler(f, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("Initialized Logging")
	return f
}
