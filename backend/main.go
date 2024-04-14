package main
 
import (
 
    "net/http"
    "sample/controller"
    "fmt"
    _"github.com/go-sql-driver/mysql"
    _"github.com/gorilla/mux"
 
)
 
 
func main() {
 
    
    // Define routes
    
 
    http.HandleFunc("/login", controller.Login)
 
    http.HandleFunc("/addPdts", controller.AddProduct)
 
    http.HandleFunc("/listPdts", controller.ListPdt)
 
    http.HandleFunc("/productDetails", controller.GetProductDetails)
 
    http.HandleFunc("/deletePdt", controller.DeletePdt)
 
    http.HandleFunc("/editPdt", controller.EditDetails)
 
 
    http.HandleFunc("/searchPdt", controller.SearchPdt)
 
     // Define the directory
     dir := "./uploads/"
 
     // Create a file server handler 
     fileServer := http.FileServer(http.Dir(dir))
  
     
  
     corsHandler := func(h http.Handler) http.Handler {
         return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
             w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
             h.ServeHTTP(w, r)
         })
     }
  
     // Register the file server handler  with CORS middleware
     http.Handle("/uploads/", corsHandler(http.StripPrefix("/uploads/", fileServer)))
  
 
 
 
 
    
 
    // Start the server
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", nil)
}