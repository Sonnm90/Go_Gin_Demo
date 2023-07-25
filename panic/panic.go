package panic

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Demo() {
	router := gin.Default()

	// Định nghĩa middleware để xử lý panic
	router.Use(RecoveryMiddleware())

	router.GET("/panic", func(c *gin.Context) {
		// Gọi hàm panic để tạo ra một lỗi panic
		panic("Panic-->Error")
	})

	_ = router.Run(":8080")

}

// RecoveryMiddleware là middleware để xử lý panic và trả về phản hồi cho người dùng
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Xử lý lỗi panic và trả về phản hồi cho người dùng
				c.JSON(500, gin.H{"error": fmt.Sprintf("Internal Server Error 123 : %v", r)})
			}
		}()
		c.Next()
	}
}
