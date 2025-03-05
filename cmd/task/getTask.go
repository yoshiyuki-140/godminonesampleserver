package task

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

func GetTask(c *gin.Context, db *gorm.DB, id int) {
	// IDがidのタスクをDBから読み込む
	var task models.Task
	err := db.First(&task, id).Error
	if err != nil {
		msg := "指定されたidのタスクは存在しません。"
		log.Println(err, "\n", msg)
		c.JSON(http.StatusNotFound, gin.H{"error": err, "message": msg})
		return
	}
	log.Println(task.Task)
	// タスクの返却
	c.JSON(http.StatusOK, task)
}
