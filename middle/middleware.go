package middle

import (
	"github.com/gin-gonic/gin"
)

func Demo() {
	// Khởi tạo router của Gin
	router := gin.Default()

	// Định nghĩa một middleware đơn giản để kiểm tra xem người dùng đã đăng nhập hay chưa
	authMiddleware := func(c *gin.Context) {
		// thực hiện xác thực người dùng
		isLoggedIn := true

		if !isLoggedIn {
			// Nếu người dùng chưa đăng nhập, trả về mã lỗi 401 và thông báo lỗi
			c.JSON(401, gin.H{"error": "Unauthorized"})
			//dừng việc thực hiện các middleware phía sau
			c.Abort()
			return
		}

		// Nếu người dùng đã đăng nhập, tiếp tục xử lý các route tiếp theo
		c.Next()
	}

	// Áp dụng middleware xác thực cho tất cả các route
	router.Use(authMiddleware)

	// Định nghĩa một route GET với đường dẫn "/hello"
	router.GET("/hello", func(c *gin.Context) {
		// Trả về chuỗi "Hello, World!" khi truy cập "/hello"
		c.String(200, "Hello, World!")
	})

	// Chạy server trên cổng 8080
	router.Run(":8080")
}
