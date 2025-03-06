package task

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

type subRequestTask struct {
	Task        string `gorm:"type:text" json:"task"`
	IsCompleted bool   `json:"is_completed"`
}

type request struct {
	SessionId string         `json:"session_id"`
	Task      subRequestTask `json:"task"`
}

// 新しいタスクを作るエンドポイント
func CreateTask(c *gin.Context, db *gorm.DB) {
	// リクエストの内容を構造体として取得する
	var req request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// セッションIDに一致するユーザーを構造体として取得
	var user models.User
	result := db.First(&user, "session_id = ?", req.SessionId)
	if err = result.Error; err != nil {
		msg := "セッションIDが無効です"
		log.Println(msg)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": msg})
		err = fmt.Errorf("%s\n%s\n", msg, err)
		return
	}

	// タスク作成
	var task models.Task
	task.UserID = int(user.ID)
	task.Task = req.Task.Task
	task.IsCompleted = req.Task.IsCompleted
	db.Create(&task)
}
