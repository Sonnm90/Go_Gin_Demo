package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Person struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"gte=18"`
}

func Demo() {
	// Khởi tạo router của Gin
	router := gin.Default()

	// Đăng ký validator với router
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("adult_age", func(fl validator.FieldLevel) bool {
			age := fl.Field().Int()
			return age >= 18
		})
	}

	// Định nghĩa một route POST với đường dẫn "/person"
	router.POST("/person", func(c *gin.Context) {
		var person Person

		// Bind dữ liệu từ yêu cầu vào struct Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		person.Name = strings.TrimSpace(person.Name)
		// Kiểm tra validation
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Thực hiện xử lý dữ liệu
		c.JSON(200, person)
	})

	// Chạy server trên cổng 8080
	router.Run(":8080")
}
