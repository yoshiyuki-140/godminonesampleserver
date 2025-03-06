package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

type deleteTaskRequest struct {
	SessionId string `json:"session_id"`
}

func DeleteTask(c *gin.Context, db *gorm.DB, id int) {
	// リクエストを取得
	var req deleteTaskRequest
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

	// タスク削除
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	db.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted", "task": task})
}
