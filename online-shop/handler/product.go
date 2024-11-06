package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"online-shop/model"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		// TODO : get data from database with ID
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("terjadi kesaalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{
				"error": "Terjadi Kesalahan pada server",
			})
			return
		}

		// TODO : give a response
		c.JSONP(200, products)
	}

}

func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO : baca id darti uri
		id := c.Param("id")

		// TODO : get data from database with ID
		product, err := model.SelectProductByID(db, id)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
				c.JSON(404, gin.H{
					"error": "product not found",
				})
				return
			}

			log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{
				"error": "Terjadi Kesalahan pada server",
			})
			return
		}

		// TODO : give a response

		c.JSONP(200, product)
	}

}

func CreateProduct(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var product model.Product

		if err := c.Bind(&product); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data product not valid"})
			return

		}

		product.ID = uuid.New().String()

		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Terjadi kesalahan saat membuat product: %v\n", err)
			c.JSON(400, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(201, product)

	}

}

func UpdateProduct(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")

		var product model.Product

		if err := c.Bind(&product); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data product not valid"})
			return

		}

		productExisting, err := model.SelectProductByID(db, id)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
				c.JSON(404, gin.H{
					"error": "product not found",
				})
				return
			}

			log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{
				"error": "Terjadi Kesalahan pada server",
			})
			return
		}

		if product.Name != "" {
			productExisting.Name = product.Name
		}

		if product.Price != 0 {
			productExisting.Price = product.Price
		}

		if err := model.UpdateProduct(db, productExisting); err != nil {
			log.Printf("Terjadi kesalahan saat update product: %v\n", err)
			c.JSON(400, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(201, productExisting)

	}

}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")

		productExisting, err := model.SelectProductByID(db, id)
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
				c.JSON(404, gin.H{
					"error": "product not found",
				})
				return
			}

			log.Printf("Terjadi Kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{
				"error": "Terjadi Kesalahan pada server",
			})
			return
		}

		if err := model.DeleteProduct(db, productExisting.ID); err != nil {
			log.Printf("Terjadi kesalahan saat update product: %v\n", err)
			c.JSON(400, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(200, "deleted")

	}

}
