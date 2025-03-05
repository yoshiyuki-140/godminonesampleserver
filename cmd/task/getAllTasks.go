package task

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

func GetAllTasks(c *gin.Context, db *gorm.DB) (err error) {
	// セッションIDを読み込み
	// タスクを読み込み
	var tasks []models.Task
	result := db.Find(&tasks)
	if err = result.Error; err != nil {
		msg := "taskを読み込めませんでした"
		err = fmt.Errorf("%s:%s\n", msg, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
	return
}
