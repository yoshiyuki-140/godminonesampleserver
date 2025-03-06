package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

// TODO: 以下の型はcreateTaskでも同じものを使っているのでInterfaceで実装したい
type updateTaskSubRequestTask struct {
	Task        string `gorm:"type:text" json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type updateTaskRequest struct {
	SessionId string         `json:"session_id"`
	Task      subRequestTask `json:"task"`
}

func UpdateTask(c *gin.Context, db *gorm.DB, id int) {

	// リクエストを取得
	var req updateTaskRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// セッションIDが有効かチェック
	var user models.User
	result := db.First(&user, "session_id = ?", req.SessionId)
	if err = result.Error; err != nil {
		msg := "セッションIDが無効です"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": msg})
		err = fmt.Errorf("%s\n%s\n", msg, err)
		return
	}

	// タスク更新
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		msg := "指定されたIDのタスクは存在しません。"
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found", "message": msg})
		return
	}
	task.Task = req.Task.Task
	task.IsCompleted = req.Task.IsCompleted
	db.Save(&task)

	c.JSON(http.StatusOK, gin.H{"message": "success", "task": task})
}
