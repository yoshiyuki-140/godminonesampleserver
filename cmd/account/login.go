package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/models"
	"gorm.io/gorm"
)

// requestの型
type request struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context, db *gorm.DB) {
	// リクエストボディーをUser構造体にバインド
	var req request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// 読み込んだデータが、存在するかどうかを判定する
	var user models.User
	result := db.First(&user, "name = ? and password = ?", req.Username, req.Password)
	if int(result.RowsAffected) == 1 {
		// TODO: セッションID生成
		session_id := "999999"
		db.Model(user).Update("session_id", session_id)

		// レスポンスでセッションIDを返却
		c.JSON(http.StatusOK, gin.H{"session_id": session_id})
	} else {
		message := "Error : ユーザーが存在しないか、パスワードが違います。"
		log.Println(message)
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
	}
	log.Println("in login func")
	return
}
