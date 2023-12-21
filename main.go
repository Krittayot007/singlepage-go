package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	ID 		string 	`json:"id"`
	Brand 	string 	`json:"brand"`
	Price 	float32 `json:"price"`
}

var cars = []Car {
	{
		ID: "1",
		Brand: "BMW",
		Price: 10000.00,
	},
	{
		ID: "2",
		Brand: "Benz",
		Price: 12000.00,
	},
	{
		ID: "3",
		Brand: "Toyota",
		Price: 5000.00,
	},
}

func main() {
	router := gin.Default()
	router.GET("/get-all-cars", getAllCar)
	router.GET("/get-cat-byID/:carsID", getCarById)
	router.POST("/add-newcar", createCar)
	router.PUT("/update-car/:carsID", updateCar)
	router.DELETE("/delete-car/:carsID", deleteCar)

	router.Run("localhost:8080")
}

func getAllCar(c *gin.Context) {
	c.JSON(http.StatusOK, cars)
}

func getCarById(c *gin.Context) {
	params := c.Params
	carsID := params[0].Value

	for _, value := range cars {
		if carsID == value.ID {
			c.JSON(http.StatusOK, value)
			return
		}	
	}
	c.JSON(http.StatusNotFound, "Not found in database")
}

func createCar(c *gin.Context) {
	var newCar Car

	if err := c.BindJSON(&newCar) ; err != nil {
		c.JSON(http.StatusBadRequest, "data not match")
		return
	}
	cars = append(cars, newCar)
	c.JSON(http.StatusCreated, "Create Successfuly")
}

func updateCar(c *gin.Context) {
	params := c.Params
	carsID := params[0].Value

	var update Car 
	if err := c.BindJSON(&update) ; err != nil {
		c.JSON(http.StatusBadRequest, "data not match")
	}

	for i := 0 ; i <= len(cars) - 1 ; i ++ {
		if carsID == cars[i].ID {
			cars[i].ID = update.ID
			cars[i].Brand = update.Brand
			cars[i].Price = update.Price 

			c.JSON(http.StatusOK, cars[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, "Data is no match")
}

func deleteCar(c *gin.Context) {
	params := c.Params
	carsID := params[0].Value

	for i := 0 ; i <= len(cars) - 1 ; i++ {
		if cars[i].ID == carsID {
			cars = append(cars[:i], cars[i+1:]...)
			c.JSON(http.StatusOK, "Delete successfully!!")
			return
		}
	}
	c.JSON(http.StatusBadRequest, "Delete incomplete")
}