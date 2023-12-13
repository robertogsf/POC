package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertogsf/POC/database"
	"github.com/robertogsf/POC/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var products []models.Product

	if err := db.Find(&products).Error; err != nil {
		http.Error(w, "Error getting products from database", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	database.DB.First(&product, params["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}
	json.NewEncoder(w).Encode(&product)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	createdProduct := database.DB.Create(&product)
	err := createdProduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	params := mux.Vars(r)
	database.DB.First(&product, params["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}
	database.DB.Delete(&product)
	w.WriteHeader(http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.DB.Model(&product).Select("*").Updates(product)
	w.Write([]byte("Successful update"))
}
