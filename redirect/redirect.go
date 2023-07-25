package redirect

import (
	"github.com/gin-gonic/gin"
)

func Demo() {
	router := gin.Default()

	// Định nghĩa middleware để thực hiện chuyển hướng từ "/old" đến "/new"
	redirectMiddleware := func(c *gin.Context) {
		if c.Request.URL.Path == "/old" {
			// Thực hiện chuyển hướng
			c.Redirect(301, "/new")
			c.Abort()
			return
		}
		c.Next()
	}

	// Áp dụng middleware chuyển hướng cho tất cả các route
	router.Use(redirectMiddleware)

	router.GET("/old", func(c *gin.Context) {
		c.String(200, "This is the old route")
	})

	router.GET("/new", func(c *gin.Context) {
		c.String(200, "This is the new route")
	})

	router.Run(":8080")
}
