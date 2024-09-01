package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {

	// 1. buat route multiplexer
	mux := http.NewServeMux()

	// 3. added handler to mux
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// 4. buat server
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 5. run server
	server.ListenAndServe()

}

var database = map[int]Products{}

var lastID = 0

// 2. fungsi handler
func listProduct(w http.ResponseWriter, r *http.Request) {

	var products []Products

	for _, v := range database {
		products = append(products, v)
	}

	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi kesalahan"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)

}

func createProduct(w http.ResponseWriter, r *http.Request) {

	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("terjadi dalam request"))
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("terjadi dalam request"))
	}

	// incremenet no uurut
	lastID++

	products.ID = lastID

	// add to database
	database[products.ID] = products

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("Request success proses and added"))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	// read ID
	productID := r.PathValue("id")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi dalam request"))
	}

	// update DB
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("terjadi dalam request"))
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("terjadi dalam request"))
	}

	products.ID = productIDInt
	// update to database
	database[productIDInt] = products
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("Request success updated"))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("terjadi dalam request"))
	}

	// delete from map
	delete(database, productIDInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
	w.Write([]byte("Request success dekleted"))

}
