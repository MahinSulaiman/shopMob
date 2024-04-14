package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	// "sample/model"
	"os"
    "path/filepath"
     "io"
	 "strconv"

	"sample/config"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	


	db := config.Connect()
	defer db.Close()
    // Parse the multipart form data
    err := r.ParseMultipartForm(10 << 20) // 10MB maximum file size
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    // Get the file from the form data
    file, handler, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Save the file to the server
    dir := "./uploads"
    os.MkdirAll(dir, os.ModePerm)
    filePath := filepath.Join(dir, handler.Filename)
    f, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer f.Close()
    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }


   name:=r.FormValue("name")
    // Generate URL for the uploaded image
  imageURL:= "http://localhost:8080/uploads/"+ handler.Filename 

   priceStr:=r.FormValue("price")
   price, err := strconv.Atoi(priceStr)
   if err != nil {
	   http.Error(w, "Invalid price value", http.StatusBadRequest)
	   return
   }
   speccs:=r.FormValue("description")

	 

    // You can now store the imageURL in your database along with other product details
	_, err = db.Exec("Insert into product(name,img_url,price,speccs) values(?,?,?,?)", name,imageURL,price,speccs)

	if err != nil {
		log.Print(err)
	}

	fmt.Println("image inserted.............")

    
   
    // w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	

	fmt.Println("image added to db....")
}
