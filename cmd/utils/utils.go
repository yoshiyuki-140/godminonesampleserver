package utils

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"gorm.io/gorm"
)

// requestの型
type request struct {
	Session_id string `gorm:"session_id"`
}

func GetUserBySessionID(c *gin.Context, db *gorm.DB, user *models.User) (err error) {
	// セッションIDを取得
	var req request
	err = c.ShouldBindJSON(&req)
	if err != nil {
		msg := "セッションIDをJSONにバインドできませんでした"
		err = fmt.Errorf("%s%s\n", msg, err)
		return
	}
	log.Println("session_id", req.Session_id)
	// セッションIDに一致するユーザーを特定する
	result := db.First(&user, "session_id = ?", req.Session_id)
	if err = result.Error; err != nil {
		msg := "セッションIDが無効です"
		log.Println(msg)
		err = fmt.Errorf("%s\n%s\n", msg, err)
	}
	return
}
