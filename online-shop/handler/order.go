package handler

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"online-shop/model"
	"time"
)

func CheckoutOrder(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var checkOutOrder model.Checkout

		// TODO : check payload request
		if err := c.Bind(&checkOutOrder); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data product not valid"})
			return
		}

		ids := []string{}
		orderQty := make(map[string]int32)
		for _, o := range checkOutOrder.Products {
			ids = append(ids, o.ID)
			orderQty[o.ID] = o.Quantity
		}

		//TODO : get data product from DB
		products, err := model.SelectProductInByIDs(db, ids)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil product detail: %v\n", err)
			c.JSON(500, gin.H{"error": "Data product not valid"})
			return
		}

		// TODO : create password
		passcode := generatePasscode(5)

		// TODO : hash password
		hashcode, err := bcrypt.GenerateFromPassword([]byte(passcode), 10)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membuat hash: %v\n", err)
			c.JSON(500, gin.H{"error": "Data product not valid"})
			return
		}
		hashcodeString := string(hashcode)

		// TODO : create order & detail
		order := model.Order{
			ID:         uuid.New().String(),
			Email:      checkOutOrder.Email,
			Address:    checkOutOrder.Address,
			GrandTotal: 0,
			Passcode:   &hashcodeString,
		}

		details := []model.OrderDetail{}

		grandTotal := int64(0)
		for _, p := range products {

			total := p.Price * int64(orderQty[p.ID])

			detail := model.OrderDetail{
				ID:        uuid.New().String(),
				OrderID:   order.ID,
				ProductID: p.ID,
				Quantity:  orderQty[p.ID],
				Price:     p.Price,
				Total:     total,
			}

			details = append(details, detail)
			grandTotal += total

		}

		order.GrandTotal = grandTotal

		model.CreateOrder(db, order, details)

		orderWithDetail := model.OrderWithDetail{
			Order:  order,
			Detail: details,
		}
		orderWithDetail.Passcode = &passcode

		c.JSON(201, orderWithDetail)
	}
}

func generatePasscode(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[randomGenerator.Intn(len(charset))]
	}

	return string(code)

}

func ConfirmOrder(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		//	TODO : get id from param
		id := c.Param("id")

		//	TODO : read request body
		var confirmReq model.Confirm
		if err := c.Bind(&confirmReq); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data order not valid"})
			return
		}

		//	TODO : get request order from db
		order, err := model.SelectOrderByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi Kesalahan saat mengambil data order: %v\n", err)
				c.JSON(404, gin.H{
					"error": "order not found",
				})
				return
			}

			log.Printf("Terjadi kesalahan saat membaca data pesanan: %v\n", err)
			c.JSON(500, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Printf("Passcode tidak valid:")
			c.JSON(500, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}

		//	TODO : compare passcode
		if err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(confirmReq.Passcode)); err != nil {
			log.Printf("terjadi kesalahan pada saat mencocokan kata sandi: %v\n:", err)
			c.JSON(401, gin.H{"error": "not access order"})
			return
		}

		//	TODO : check must be order not paid
		if order.PaidAt != nil {
			log.Printf("order has been paid: %v\n:")
			c.JSON(401, gin.H{"error": "order has been paid"})
			return
		}

		//	TODO : compare amount send from amount from order
		if order.GrandTotal != confirmReq.Amount {
			log.Printf("jumlah harga tidak sesuai: %d\n", confirmReq.Amount)
			c.JSON(401, gin.H{"error": "total payment is not valid"})
			return
		}

		//	TODO : update data order
		currentTimeNow := time.Now()
		if err = model.UpdateOrderByID(db, id, confirmReq, currentTimeNow); err != nil {
			log.Printf("Error on update data order: %v\n", err)
			c.JSON(500, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}

		order.PaidAccountNumber = &confirmReq.AccountNumber
		order.PaidAt = &currentTimeNow
		order.PaidBank = &confirmReq.Bank
		order.Passcode = &confirmReq.Passcode

		c.JSON(201, order)

	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		//	TODO : get id from param
		id := c.Param("id")

		//TODO : ambil passcode dari query param
		passcode := c.Query("passcode")

		//	TODO : read request body
		var confirmReq model.Confirm
		if err := c.Bind(&confirmReq); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data order not valid"})
			return
		}

		//	TODO : get request order from db
		order, err := model.SelectOrderByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi Kesalahan saat mengambil data order: %v\n", err)
				c.JSON(404, gin.H{
					"error": "order not found",
				})
				return
			}

			log.Printf("Terjadi kesalahan saat membaca data pesanan: %v\n", err)
			c.JSON(500, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Printf("Passcode tidak valid:")
			c.JSON(500, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}

		//	TODO : compare passcode
		if err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(passcode)); err != nil {
			log.Printf("terjadi kesalahan pada saat mencocokan kata sandi: %v\n:", err)
			c.JSON(401, gin.H{"error": "not access order"})
			return
		}

		order.Passcode = nil

		c.JSON(200, order)

	}
}
