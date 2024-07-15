package web

import (
	"database/sql"
	"go_task_list/internal/models"
	database "go_task_list/internal/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

var helper database.DB_Helper

func StartRouter(db *sql.DB) {
	helper.Db = db
	router := gin.Default()
	router.Static("/static", "../web/static")
	router.LoadHTMLFiles("../web/templates/default.html", "../web/templates/404.html")
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Task list",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default.html", gin.H{
			"title": "Task list",
		})
	})

	router.GET("/get_tasks", func(c *gin.Context) {
		tasks, err := helper.GetAllTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error during fetching tasks from db"})
		} else {
			c.JSON(http.StatusOK, tasks)
		}
	})

	router.POST("/insert_task", func(c *gin.Context) {
		task := models.Task{}
		c.BindJSON(&task)
		err := helper.InsertTask(task.User_id, task.Created_at, task.Title, task.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error during inserting tasks to db"})
		} else {
			c.JSON(http.StatusOK, "Successful")
		}
	})

	router.Run(":8080")
}
