package controller

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"sample/model"

	// "os"
	// "path/filepath"
	//  "io"
	"strconv"

	"sample/config"
)

func GetProductDetails(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Get the product ID from the query parameter
	idStr := r.URL.Query().Get("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Find the product with the given ID
	db := config.Connect()
	defer db.Close()

	err = db.QueryRow("SELECT img_url,speccs,price,name FROM product WHERE id=?", productID).Scan(&product.ImageURL, &product.Speccs, &product.Price, &product.Name)
	if err != nil {
		fmt.Println("error in query", err)
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If product not found, return 404
	// if product == nil {
	// 	http.Error(w, "Product not found", http.StatusNotFound)
	// 	return
	// }

	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("error in marshal data")
		return
	}

	// fmt.Println(jsonData)

	fmt.Println("viewing one......")

	w.Write(jsonData)

}
