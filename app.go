package main

import (
	"changeme/app/handler/menu"
	"changeme/app/handler/note"
	"changeme/app/middleware"
	"changeme/pkg/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// StartHttp http service
func StartHttp() {
	// db open
	model.Conn()
	r := gin.Default()
	r.Use(middleware.Cors())
	menu_handler := new(menu.MenuHandler)
	note_handler := new(note.NoteHandler)
	gin.SetMode("debug")

	r.GET("/api/menu", menu_handler.Index)
	r.POST("/api/menu", menu_handler.Store)
	r.DELETE("/api/menu/:id", menu_handler.Delete)
	r.PUT("/api/update", menu_handler.Update)

	r.GET("/api/note/:id", note_handler.Index)
	r.POST("/api/note", note_handler.Store)
	r.DELETE("/api/note/:id", note_handler.Delete)
	r.PUT("/api/note/:id", note_handler.Update)

	r.Run(":7999") // listen and serve on 0.0.0.0:8080
}
