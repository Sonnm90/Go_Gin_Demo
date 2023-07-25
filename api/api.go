package api

import "github.com/gin-gonic/gin"

func Demo() {
	// Khởi tạo router của Gin
	router := gin.Default()

	// Định nghĩa một route GET với đường dẫn "/hello"
	router.GET("/hello", func(c *gin.Context) {
		// Trả về chuỗi "Hello, World!" khi truy cập "/hello"
		c.String(200, "Hello, World!")
	})

	// Chạy server trên cổng 8080
	_ = router.Run(":8080")
}
