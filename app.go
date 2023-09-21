package main

import (
	"changeme/pkg/db"
	"changeme/pkg/helpers"
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

func StartHttp() {
	// db open
	db.Conn()
	r := gin.Default()

	r.GET("/api/menu", func(c *gin.Context) {
		rows, err := db.DB.Query("select * from note_menus where is_deleted = 0")
		helpers.CheckErr(err)
		var menuList = []MenuList{}
		fmt.Println(rows.Next())
		for rows.Next() {
			var menu MenuList
			err = rows.Scan(&menu.MenuId, &menu.Name, &menu.PId,
				&menu.AddTime, &menu.IsDeleted, &menu.IsDir)
			menuList = append(menuList, menu)
		}
		c.JSON(200, GetMenuTree(menuList, 0))

	})
	r.POST("/api/menu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
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

type MenuList struct {
	MenuId    int    `json:"menu_id"`
	Name      string `json:"name"`
	PId       int    `json:"p_id"`
	AddTime   int    `json:"add_time"`
	IsDeleted int    `json:"is_deleted"`
	IsDir     int    `json:"is_dir"`
	Children  any    `json:"children,omitempty"`
}

func GetMenuTree(menus []MenuList, pid int) any {
	var list []MenuList
	for _, v := range menus {
		if v.PId == pid {
			v.Children = GetMenuTree(menus, v.MenuId)
			list = append(list, v)
		}
	}
	return list
}
