package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	// "sample/model"
	"os"
    // "path/filepath"
    //  "io"
	 "strconv"

	"sample/config"
)

func DeletePdt(w http.ResponseWriter,r *http.Request){
	var imgUrl string
	

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
		
		db:=config.Connect()
		defer db.Close()

		err=db.QueryRow("select img_url from product where id=?",productID).Scan(&imgUrl)
		if err != nil {
			log.Print(err)
		}
		imageName:= path.Base(imgUrl)




		// deleting from the folder
		imagePath := "./uploads/" + imageName 
		err = os.Remove(imagePath)
		if err != nil {
			fmt.Println("failed to delete  error ", err)
			http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
			return
		}
		

		_, err = db.Exec("delete  from product where id=?",productID)

		if err != nil {
			log.Print(err)
		}

		
		
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Println("deleted")



}