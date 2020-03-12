package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	router.GET("/product", getProduct)
	router.GET("/product/:category", getProductCategory)

	router.Run()
}

type Data struct {
	ID 						int				`json:"id"`
	ProductName 	string		`json:"product_name"`
	Image 				string		`json:"image"`
	Category 			string		`json:"category"`
	Price 				int				`json:"price"`
}

type Response struct {
	Status 	string	`json:"status"`
	Message string	`json:"message"`
	Data 		[]Data `json:"data"`
}

func getProduct(context *gin.Context) {
	responseData := Response{
		Status	: "success",
		Message	: "data has been loaded",
		Data		:	[]Data{
			{
				ID						: 1,
				ProductName		:	"Buku Membangun Aplikasi Android Web Dan Web Service",
				Image					: "https://cdn.gramedia.com/uploads/items/membangun_aplikasi__w414_hauto.jpg",
				Category			: "technology",
				Price					: 85000,
			},
			{
				ID						: 2,
				ProductName		:	"Buku Menguasai Pemrograman Berorientasi Objek",
				Image					: "https://cdn.gramedia.com/uploads/items/menguasai_pemrograman__w414_hauto.jpg",
				Category			: "technology",
				Price					: 125000,
			},
			{
				ID						: 3,
				ProductName		:	"Buku Fungsi Variabel Kompleks",
				Image					: "https://cdn.gramedia.com/uploads/items/fungsi_variable__w414_hauto.jpg",
				Category			: "education",
				Price					: 88000,
			},
			{
				ID						: 4,
				ProductName		:	"Buku Teach Like Fun Teacher",
				Image					: "https://cdn.gramedia.com/uploads/items/img293__w414_hauto.jpg",
				Category			: "education",
				Price					: 57000,
			},
			{
				ID						: 5,
				ProductName		:	"Buku Evaluasi Pembelajaran",
				Image					: "https://cdn.gramedia.com/uploads/items/img200__w414_hauto.jpg",
				Category			: "education",
				Price					: 60000,
			},
		},
	}
	context.JSON(200, responseData)
}

func getProductCategory(context *gin.Context) {
	category := context.Param("category")
	if category == "technology" {
		responseData := Response{
			Status	: "success",
			Message	: "data has been loaded",
			Data		:	[]Data{
				{
					ID						: 1,
					ProductName		:	"Buku Membangun Aplikasi Android Web Dan Web Service",
					Image					: "https://cdn.gramedia.com/uploads/items/membangun_aplikasi__w414_hauto.jpg",
					Category			: "technology",
					Price					: 85000,
				},
				{
					ID						: 2,
					ProductName		:	"Buku Menguasai Pemrograman Berorientasi Objek",
					Image					: "https://cdn.gramedia.com/uploads/items/menguasai_pemrograman__w414_hauto.jpg",
					Category			: "technology",
					Price					: 125000,
				},
			},
		}
		context.JSON(200, responseData)
	} else if category == "education" {
		responseData := Response{
			Status	: "success",
			Message	: "data has been loaded",
			Data		:	[]Data{
				{
					ID						: 3,
					ProductName		:	"Buku Fungsi Variabel Kompleks",
					Image					: "https://cdn.gramedia.com/uploads/items/fungsi_variable__w414_hauto.jpg",
					Category			: "education",
					Price					: 88000,
				},
				{
					ID						: 4,
					ProductName		:	"Buku Teach Like Fun Teacher",
					Image					: "https://cdn.gramedia.com/uploads/items/img293__w414_hauto.jpg",
					Category			: "education",
					Price					: 57000,
				},
				{
					ID						: 5,
					ProductName		:	"Buku Evaluasi Pembelajaran",
					Image					: "https://cdn.gramedia.com/uploads/items/img200__w414_hauto.jpg",
					Category			: "education",
					Price					: 60000,
				},
			},
		}
		context.JSON(200, responseData)
	} else {
		context.JSON(200, gin.H {
			"status"	: "success",
			"message"	: "data has not found",
		})
	}
}
