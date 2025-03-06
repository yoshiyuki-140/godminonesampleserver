package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/account"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/models"
	"github.com/yoshiyuki-140/godminonesampleserver/cmd/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := "localhost" // または環境変数から取得
	dbPort := "5432"      // または環境変数から取得

	// DSNを構築
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	// GORMでデータベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// データベースにテーブルを作成
	db.AutoMigrate(&models.User{}, &models.Task{})

	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// テストエントリポイント
	r.GET("/", func(c *gin.Context) {
		msg := "Hello World"
		c.JSON(200, gin.H{"message": msg})
	})

	// register
	r.POST("/account/register", func(c *gin.Context) {
		account.Register(c, db)
	})

	// login
	r.GET("/account/login", func(c *gin.Context) {
		account.Login(c, db)
		log.Println("Hello")
	})

	// logout
	r.POST("/account/logout", func(c *gin.Context) {
		account.Logout(c, db)
	})

	// 全てのタスク一覧を取得する
	r.GET("/tasks", func(c *gin.Context) {
		task.GetAllTasks(c, db)
	})

	// タスクを一つ取得する
	r.GET("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "パスパラメータを読み込めませんでした。\nこのエントリポイントでのパスパラメータは数値にしてください。Example : localhost:8080/tasks/1"})
			return
		}
		task.GetTask(c, db, id)
	})

	// 新しいタスクを作成するエンドポイント
	r.POST("/tasks", func(c *gin.Context) {
		task.CreateTask(c, db)
	})

	// タスクを更新するエンドポイント
	r.PUT("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "パスパラメータを読み込めませんでした。",
				"message": "このエントリポイントでのパスパラメータは数値にしてください。",
				"example": "localhost:8080/tasks/1",
			})
			return
		}
		task.UpdateTask(c, db, id)
	})

	/* TODO: タスク削除のハンドラ関数を別ファイルに分割 */
	// タスクを削除するエンドポイント
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task models.Task
		id := c.Param("id")

		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		db.Delete(&task)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
