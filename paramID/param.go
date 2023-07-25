package paramID

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Product 1", Price: 100},
	{ID: 2, Name: "Product 2", Price: 200},
	{ID: 3, Name: "Product 3", Price: 300},
}

func Demo() {
	router := gin.Default()

	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)

	router.Run(":8080")
}

func getProducts(c *gin.Context) {
	c.JSON(200, products)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, p := range products {
		if strconv.Itoa(p.ID) == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Product not found"})
}
