package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"io"
	"os"
	"path/filepath"
	"strconv"

	"path"

	"sample/config"
)

func EditDetails(w http.ResponseWriter, r *http.Request) {
	var imgUrl string
	fmt.Println("edit portion")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	

	db := config.Connect()
	defer db.Close()
	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10MB maximum file size
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("image")
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

    //delete existing file from folder

	err=db.QueryRow("select img_url from product where id=?",productID).Scan(&imgUrl)
		if err != nil {
			log.Print(err)
		}
		imageName:= path.Base(imgUrl)

	imagePath := "./uploads/" + imageName
		err = os.Remove(imagePath)
		if err != nil {
			fmt.Println("failed to delete  error ", err)
			http.Error(w, "Failed to delete image file", http.StatusInternalServerError)
			return
		}

	// Save the file to the server
	dir := "./uploads"
	os.MkdirAll(dir, os.ModePerm)
	filePath := filepath.Join(dir, handler.Filename)
	f, err := os.Create(filePath)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")

	// Generate URL for the uploaded image
	imageURL := "http://localhost:8080/uploads/" + handler.Filename

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}
	speccs := r.FormValue("speccs")

	fmt.Println(speccs)

	fmt.Println(productID)
	// update query
	_, err = db.Exec("UPDATE product SET img_url=?, speccs=? , price=?,name=? WHERE id=?",   imageURL,speccs, price,name,productID)
	if err != nil {
		fmt.Println("eror in db")
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}

	// fmt.Println("image updated.............")

	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})


	// fmt.Println("image added again to db....")

}
