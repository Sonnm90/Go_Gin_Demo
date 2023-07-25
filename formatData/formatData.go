package formatData

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func Demo() {
	// Khởi tạo router của Gin
	router := gin.Default()

	// Định nghĩa một route GET với đường dẫn "/json"
	router.GET("/json", func(c *gin.Context) {
		// Trả về dữ liệu dưới định dạng JSON
		c.JSON(200, Person{Name: "John", Age: 30})
	})

	// Định nghĩa một route GET với đường dẫn "/xml"
	router.GET("/xml", func(c *gin.Context) {
		// Trả về dữ liệu dưới định dạng XML
		c.XML(200, Person{Name: "Jane", Age: 25})
	})
	router.GET("/yaml", func(c *gin.Context) {
		// Trả về dữ liệu dưới định dạng XML
		c.YAML(200, Person{Name: "Jane", Age: 25})
	})
	router.GET("/toml", func(c *gin.Context) {
		// Trả về dữ liệu dưới định dạng XML
		c.TOML(200, Person{Name: "Jane", Age: 25})
	})

	// Chạy server trên cổng 8080
	router.Run(":8080")
}
