package routers

import (
	"encoding/json"
	"net/http"

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
