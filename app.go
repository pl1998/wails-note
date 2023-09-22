package main

import (
	"changeme/app/handler"
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

	menu_handler := new(handler.MenuHandler)

	r.GET("/api/menu", menu_handler.Index)
	r.POST("/api/menu", menu_handler.Store)
	r.DELETE("/api/menu/:id", menu_handler.Delete)
	r.PUT("/api/update", menu_handler.Update)

	r.DELETE("/api/menu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/note", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.PUT("/api/note", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.DELETE("/api/note", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":7999") // listen and serve on 0.0.0.0:8080
}
