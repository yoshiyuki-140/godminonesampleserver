package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/utils"
	"gorm.io/gorm"
)

func Logout(c *gin.Context, db *gorm.DB) {
	// セッションIDからユーザーを取得
	var user models.User
	err := utils.GetUserBySessionID(c, db, &user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// そのユーザーの持つセッションIDを0にリセットする
	log.Println("user:", user)
	user.SessionId = 0
	db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Success", "msg": user.ID})
	return
}
