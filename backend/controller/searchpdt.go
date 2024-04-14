package controller

import (
	"encoding/json"
	"fmt"
	
	"net/http"
	"sample/model"
	

	"sample/config"
)

func SearchPdt(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	fmt.Println(query)
	if query == "" {
		fmt.Println("empty")
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	db:=config.Connect()
		defer db.Close()
		
	rows, err := db.Query("SELECT id, name, speccs, img_url FROM product WHERE name LIKE ? OR speccs LIKE ?", "%"+query+"%", "%"+query+"%")
	if err != nil {
		http.Error(w, "Error searching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var arrPdt  []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Speccs, &product.ImageURL)
		if err != nil {
			http.Error(w, "Error scanning products", http.StatusInternalServerError)
			return
		}
		arrPdt = append(arrPdt, product)
	}
	
	json.NewEncoder(w).Encode(arrPdt)
	fmt.Println("searching")

}
