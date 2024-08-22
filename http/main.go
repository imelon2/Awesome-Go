package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

var users = []User{}
var nextID = 1

func main() {
	r := gin.Default()

	// GET 요청: 모든 사용자 조회
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// POST 요청: 새로운 사용자 생성
	r.POST("/user", func(c *gin.Context) {
		var user User

		// JSON 바인딩
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 사용자 ID 설정 및 저장
		user.ID = nextID
		nextID++
		users = append(users, user)

		c.JSON(http.StatusOK, gin.H{
			"status": "User created successfully",
			"user":   user,
		})
	})

	r.Run(":8080") // 서버를 8080 포트에서 실행
}
