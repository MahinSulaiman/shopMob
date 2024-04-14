package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sample/model"
	

	"sample/config"
)

func ListPdt(w http.ResponseWriter, r *http.Request){

	
	 
	var arrPdt []model.Product
	var product model.Product

	db := config.Connect()
	defer db.Close()

	rows,err:=db.Query("select id,img_url,name,price from product")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err := rows.Scan(&product.ID,&product.ImageURL,&product.Name,&product.Price)
		

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrPdt = append(arrPdt,product)
		}
	}

	
	jsonData,err:=json.Marshal(arrPdt)
	if err!=nil{
		fmt.Println("error in marshal data")
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(jsonData)

	fmt.Println("viewing.........")

	w.Write(jsonData)

   
}