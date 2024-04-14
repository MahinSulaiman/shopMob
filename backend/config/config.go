package config
import "database/sql"

import "fmt"

func Connect()*sql.DB{
	db, err := sql.Open("mysql", "mahin:password@tcp(127.0.0.1:3306)/sample_db")
    // db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sample_db")

    
    if err != nil {
        fmt.Println("error db")
        panic(err.Error())
       
    }
	return db
}

