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
	//  "strconv"

	"sample/config"
)

func Login(w http.ResponseWriter,r *http.Request){

	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return
    }
    // Set CORS headers for the main request
    w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")

	w.Header().Set("Content-Type", "application/json")


	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
	db := config.Connect()
	defer db.Close()

	err = db.QueryRow("SELECT role FROM user WHERE name = ? AND password = ?", user.Username, user.Password).Scan(&user.Role)
	if err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	
 
	w.WriteHeader(http.StatusOK)

    w.Write(jsonData)

}











